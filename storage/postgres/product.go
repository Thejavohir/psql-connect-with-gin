package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"psql/api/models"
	"psql/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO product(id, name, price, category_id, barcode, updated_at)
			VALUES ($1, $2, $3, $4, $5, NOW())`

	_, err := p.db.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		req.CategoryID,
		req.Barcode,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *productRepo) GetById(ctx context.Context, req *models.ProductPKey) (*models.Product, error) {

	var (
		categoryObj pgtype.JSON
		branchObj   pgtype.JSON

		id         sql.NullString
		name       sql.NullString
		price      sql.NullFloat64
		categoryId sql.NullString
		barcode    sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		WITH product_branches AS (
		    SELECT
		        bp.product_id,
		        JSON_AGG(
		            JSON_BUILD_OBJECT(
		                'id', b.id,
		                'name', b.name,
		                'address', b.address,
		                'phone_number', b.phone_number,
		                'created_at', b.created_at,
		                'updated_at', b.updated_at
		            )
		        ) AS branches
		    FROM branch AS b
		    JOIN branch_product_relation AS bp ON bp.branch_id = b.id
		    WHERE bp.product_id = $1
		    GROUP BY bp.product_id
		)
		SELECT
		    p.id,
		    p.name,
		    p.price,
		    p.category_id,
			p.barcode,
		    p.created_at,
		    p.updated_at,
		    COALESCE(pb.branches, '[]'::json) as branches,

		    JSON_BUILD_OBJECT(
		    'id', c.id,
		    'title', c.title,
		    'parent_id', c.parent_id,
		    'updated_at', c.updated_at,
		    'created_at', c.created_at
		    ) AS category
		FROM product AS p 
		JOIN product_branches AS pb ON pb.product_id = p.id
		LEFT JOIN category AS c ON c.id = p.category_id
		WHERE p.id = $1;
	`
	err := r.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&name,
		&price,
		&categoryId,
		&barcode,
		&createdAt,
		&updatedAt,
		&branchObj,
		&categoryObj,
	)
	if err != nil {
		return nil, err
	}

	branch := []*models.Branch{}
	err = branchObj.AssignTo(&branch)
	if err != nil {
		return nil, fmt.Errorf("assigning branch: %w", err)
	}

	category := models.Category{}
	err = categoryObj.AssignTo(&category)
	if err != nil {
		return nil, fmt.Errorf("assigning category: %w", err)
	}

	return &models.Product{
		ID:           id.String,
		Name:         name.String,
		Price:        price.Float64,
		CategoryData: &category,
		Branches:     branch,
		CategoryID:   categoryId.String,
		Barcode:      barcode.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.ProductGetListReq) (*models.ProductGetListResp, error) {

	var (
		resp = &models.ProductGetListResp{}
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
			name,
			price,
			category_id,
			barcode,
			created_at,
			updated_at
		FROM product
		`
	// JSON_BUILD_OBJECT(
	// 'id', c.id,
	// 'title', c.title,
	// 'parent_id', c.parent_id,
	// 'created_at', c.created_at,
	// 'updated_at', c.updated_at
	// ) AS category
	// LEFT JOIN category AS c ON c.id = p.category_id

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	if req.Barcode != "" {
		where += ` AND barcode = ` + req.Barcode
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	} 
	defer rows.Close()

	for rows.Next() {
		var (
			id         sql.NullString
			name       sql.NullString
			price      sql.NullFloat64
			categoryId sql.NullString
			barcode    sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&categoryId,
			&barcode,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Product = append(resp.Product, &models.Product{
			ID:         id.String,
			Name:       name.String,
			Price:      price.Float64,
			CategoryID: categoryId.String,
			Barcode:    barcode.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {

	var (
		params map[string]interface{}
	)

	query := `
		UPDATE product SET
		name = :name,
		price = :price,
		category_id = :category_id,
		barcode = :barcode,
		updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":          req.ID,
		"name":        req.Name,
		"price":       req.Price,
		"category_id": helper.NewNullString(req.CategoryID),
		"barcode": req.Barcode,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	resp, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return resp.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPKey) error {
	query := `
		DELETE FROM product
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
		UPDATE product SET ` + set + ` updated_at = NOW()
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
