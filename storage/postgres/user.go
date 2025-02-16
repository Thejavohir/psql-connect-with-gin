package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"psql/api/models"
	"psql/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (p *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO users(id, username, password, updated_at)
			VALUES ($1, $2, $3, NOW())`

	_, err := p.db.Exec(ctx, query,
		id,
		req.Username,
		req.Password,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetById(ctx context.Context, req *models.UserPKey) (*models.User, error) {

	var idOrUsername = "id"
	if len(req.Username) > 0 {
		idOrUsername = "username"
		req.ID = req.Username
	}

	var (
		id        sql.NullString
		username  sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `SELECT
				id,
				username,
				password,
				created_at,
				updated_at	
			FROM users
			WHERE ` + idOrUsername + ` = $1 `
	err := r.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&username,
		&password,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        id.String,
		Username:  username.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, req *models.UserPKey) (*models.User, error) {

	var (
		id        sql.NullString
		username  sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `SELECT
				id,
				username,
				password,
				created_at,
				updated_at	
			FROM users
			WHERE username = $1`
	err := r.db.QueryRow(ctx, query, req.Username).Scan(
		&id,
		&username,
		&password,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        id.String,
		Username:  username.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.UserGetListReq) (*models.UserGetListResp, error) {

	var (
		resp = &models.UserGetListResp{}
		// categoryObj pgtype.JSON
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 0"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			username,
			password,
			created_at,
			updated_at
		FROM users
		`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        sql.NullString
			username  sql.NullString
			password  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&username,
			&password,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Users = append(resp.Users, &models.User{
			ID:        id.String,
			Username:  username.String,
			Password:  password.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		params map[string]interface{}
	)

	query := `
		UPDATE users SET
		name = :name,
		price = :price,
		category_id = :category_id,
		barcode = :barcode,
		updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":       req.ID,
		"username": req.Username,
		"password": req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	resp, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return resp.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPKey) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE users SET ` + set + ` updated_at = NOW()
		WHERE id = :id`

	req.Fields["id"] = req.ID

	query, args := helper.ReplaceQueryParams(query, req.Fields)
	resp, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return resp.RowsAffected(), nil
}
