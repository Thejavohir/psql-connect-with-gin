package postgres

import (
	"database/sql"
	"fmt"
	"psql/models"
	"psql/pkg/helper"

	"github.com/google/uuid"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(req *models.CreateCategory) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO category(id, title, parent_id, updated_at)
			VALUES($1, $2, $3, NOW())`

	_, err := r.db.Exec(query,
		id,
		req.Title,
		helper.NewNullString(req.ParentID),
	)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (r *categoryRepo) GetById(req *models.CategoryPKey) (*models.Category, error) {

	var (
		resp  models.Category
	)

	query := `
		SELECT
			id,
			title,
			COALESCE(parent_id::VARCHAR, ''),
			created_at,
			updated_at
		FROM category WHERE id = $1
	`
	err := r.db.QueryRow(query, req.ID).Scan(
		&resp.ID,
		&resp.Title,
		&resp.ParentID,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *categoryRepo) GetList(req *models.CategoryGetListReq) (*models.CategoryGetListResp, error) {

	var (
		resp = &models.CategoryGetListResp{}
		query string
		where = " WHERE TRUE"
		offset = " OFFSET 0"
		limit = " LIMIT 0"
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

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&resp.Count,
			&category.ID,
			&category.Title,
			&category.ParentID,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Category = append(resp.Category, &category)
	}

	return resp, nil
}

func (r *categoryRepo) Update(req *models.UpdateCategory) (*models.Category, error) {

	var (
		resp models.Category
	)

	query := `
		UPDATE category SET
		title = $1,
		parent_id = $2,
		updated_at = NOW()
		WHERE id = $3
		RETURNING id, title, COALESCE(parent_id::VARCHAR, ''), created_at, updated_at
	`
	err := r.db.QueryRow(query, req.Title, req.ParentID, req.ID).Scan(
		&resp.ID,
		&resp.Title,
		&resp.ParentID,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *categoryRepo) Delete(req *models.CategoryPKey) error {
	query := `
		DELETE FROM category
		WHERE id = $1
	`

	_, err := r.db.Exec(query, req.ID)
	if err != nil {
		return nil
	}

	return nil
}