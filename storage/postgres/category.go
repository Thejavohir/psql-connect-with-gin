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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(pgxpool *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: pgxpool,
	}
}

func (r *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO category(id, title, parent_id, updated_at)
			VALUES($1, $2, $3, NOW())`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Title,
		helper.NewNullString(req.ParentID),
	)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (r *categoryRepo) GetById(ctx context.Context, req *models.CategoryPKey) (*models.Category, error) {

	var (
		id        sql.NullString
		title     sql.NullString
		parentId  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			id,
			title,
			parent_id,
			created_at,
			updated_at
		FROM category WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&title,
		&parentId,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:        id.String,
		Title:     title.String,
		ParentID:  parentId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *categoryRepo) GetList(ctx context.Context, req *models.CategoryGetListReq) (*models.CategoryGetListResp, error) {

	var (
		resp   = &models.CategoryGetListResp{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			title,
			COALESCE(parent_id::VARCHAR, ''),
			created_at,
			updated_at
		FROM category
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
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
			title     sql.NullString
			parentId  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&parentId,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Category = append(resp.Category, &models.Category{
			ID:        id.String,
			Title:     title.String,
			ParentID:  parentId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		params map[string]interface{}
	)

	query := `
		UPDATE category SET
		title = :title,
		parent_id = :parent_id,
		updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":        req.ID,
		"title":     req.Title,
		"parent_id": helper.NewNullString(req.ParentID),
	}

	query, args := helper.ReplaceQueryParams(query, params)
	resp, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return resp.RowsAffected(), nil
}

func (r *categoryRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {
	var (
		query string
		set   string
	)

	if len(req.Fields) <= 0 {
		return 0, errors.New("no fields")
	}

	for key := range req.Fields {
		fmt.Println("KEEYYY:", key)
		set += fmt.Sprintf(" %s = :%s, ", key, key)
	}

	query = `
		UPDATE category SET ` + set + ` updated_at = NOW()
		WHERE id = :id`

	req.Fields["id"] = req.ID

	fmt.Println("QUERYYY:", query)

	query, args := helper.ReplaceQueryParams(query, req.Fields)
	resp, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return resp.RowsAffected(), nil
}

func (r *categoryRepo) Delete(ctx context.Context, req *models.CategoryPKey) error {

	_, err := r.db.Exec(ctx, `DELETE FROM category WHERE id = $1`, req.ID)
	if err != nil {
		return err
	}

	return nil
}
