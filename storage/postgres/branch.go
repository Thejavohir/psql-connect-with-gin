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
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (p *branchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {

	trx, err := p.db.Begin(ctx)
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		id    = uuid.New().String()
		query string
	)

	query = `INSERT INTO branch(id, name, address, phone_number, updated_at)
			VALUES ($1, $2, $3, $4, NOW())`

	_, err = trx.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		return "", err
	}

	// if len(req.ProductIDs) > 0 {
	// 	branchProductInsert := `INSERT INTO branch_product_relation(branch_id, product_id) VALUES `
	// 	branchProductInsert, args := helper.InsertMultiple(branchProductInsert, id, req.ProductIDs)
	// 	_, err = trx.Exec(ctx, branchProductInsert, args...)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }

	return id, nil
}

func (r *branchRepo) GetById(ctx context.Context, req *models.BranchPKey) (*models.Branch, error) {

	var (
		id          sql.NullString
		name        sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString

		productsObj pgtype.JSON
	)

	query := `
		WITH branch_products AS (
			SELECT
				mp.branch_id,
				JSON_AGG(
					JSONB_BUILD_OBJECT(
						'id', p.id,
						'name', p.name,
						'price', p.price,
						'category_id', p.category_id,
						'created_at', p.created_at,
						'updated_at', p.updated_at
					)
				) AS products
			FROM product AS p
			JOIN branch_product_relation AS mp ON mp.product_id = p.id
			WHERE mp.branch_id = $1
			GROUP BY mp.branch_id
		)
		SELECT 
			m.id,
			m.name,
			m.address,                                        
			m.phone_number,
			m.created_at,
			m.updated_at ,
			COALESCE(mp.products, '[]'::json) AS products
		FROM branch AS m
		JOIN branch_products AS mp ON mp.branch_id = m.id 
		WHERE m.id = $1;
	`
	err := r.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&name,
		&address,
		&phoneNumber,
		&createdAt,
		&updatedAt,
		&productsObj,
	)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	if len(productsObj.Bytes) > 0 {
		err = productsObj.AssignTo(&products)
		if err != nil {
			return nil, fmt.Errorf("JSON parsing xatosi: %v", err)
		}
	}	

	return &models.Branch{
		ID:          id.String,
		Name:        name.String,
		Address:     address.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
		Products:    products,
	}, nil
}

func (r *branchRepo) GetList(ctx context.Context, req *models.BranchGetListReq) (*models.BranchGetListResp, error) {

	var (
		resp   = &models.BranchGetListResp{}
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
			address,
			phone_number,
			created_at,
			updated_at
		FROM branch `

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
			id           sql.NullString
			name         sql.NullString
			address      sql.NullString
			phone_number sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&phone_number,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Branch = append(resp.Branch, &models.Branch{
			ID:          id.String,
			Name:        name.String,
			Address:     address.String,
			PhoneNumber: phone_number.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *branchRepo) Update(ctx context.Context, req *models.UpdateBranch) (int64, error) {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer func(trx pgx.Tx) {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}(trx)

	var (
		params map[string]interface{}
	)

	query := `
		UPDATE branch SET
		name = :name,
		address = :address,
		phone_number = :phone_number,
		updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.ID,
		"name":         req.Name,
		"address":      req.Address,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	resp, err := trx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	existingProductIDs := make(map[string]interface{})
	rows, err := trx.Query(ctx, `SELECT product_id FROM branch_product_relation WHERE branch_id = $1`, req.ID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var productID string
		if err := rows.Scan(&productID); err != nil {
			return 0, err
		}
		existingProductIDs[productID] = struct{}{}
	}

	var newProducts []string
	for _, producID := range req.Products {
		if _, exists := existingProductIDs[producID]; !exists {
			newProducts = append(newProducts, producID)
		}
	}

	if len(newProducts) > 0 {
		branchProductInsert := `INSERT INTO branch_product_relation(branch_id, product_id) VALUES `
		branchProductInsert, args := helper.InsertMultiple(branchProductInsert, req.ID, newProducts)
		_, err = trx.Exec(ctx, branchProductInsert, args...)
		if err != nil {
			return 0, err
		}
	}

	return resp.RowsAffected(), nil
}

func (r *branchRepo) Delete(ctx context.Context, req *models.BranchPKey) error {
	query := `
		DELETE FROM branch
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *branchRepo) Patch(ctx context.Context, req *models.PatchRequest) (int64, error) {

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
		UPDATE branch SET ` + set + ` updated_at = NOW()
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
