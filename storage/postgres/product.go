package postgres

import (
	"database/sql"
	"fmt"
	"psql/models"

	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(req *models.CreateProduct) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO product(id, name, price, category_id, updated_at)
			VALUES ($1, $2, $3, $4, NOW())`

	_, err := p.db.Exec(query,
		id,
		req.Name,
		req.Price,
		req.CategoryID,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *productRepo) GetById(req *models.ProductPKey) (*models.Product, error) {

	var (
		resp models.Product
	)

	query := `
		SELECT
			id,
			name,
			price,
			category_id,
			created_at,
			updated_at
		FROM product WHERE id = $1
	`
	err := r.db.QueryRow(query, req.ID).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Price,
		&resp.CategoryID,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *productRepo) GetList(req *models.ProductGetListReq) (*models.ProductGetListResp, error) {

	var (
		resp   = &models.ProductGetListResp{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 0"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			category_id,
			created_at,
			updated_at
		FROM product
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
		var product models.Product
		err := rows.Scan(
			&resp.Count,
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Product = append(resp.Product, &product)
	}

	return resp, nil
}

func (r *productRepo) Update(req *models.UpdateProduct) (*models.Product, error) {

	var (
		resp models.Product
	)

	query := `
		UPDATE product SET
		name = $1,
		price = $2,
		category_id = $3,
		updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, price, category_id, created_at, updated_at
	`
	err := r.db.QueryRow(query, req.Name, req.Price, req.CategoryID, req.ID).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Price,
		&resp.CategoryID,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *productRepo) Delete(req *models.ProductPKey) error {
	query := `
		DELETE FROM product
		WHERE id = $1
	`

	_, err := r.db.Exec(query, req.ID)
	if err != nil {
		return nil
	}

	return nil
}
