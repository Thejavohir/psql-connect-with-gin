package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"psql/config"
	"psql/storage"
)

type store struct {
	db *sql.DB
	category *categoryRepo
	product *productRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	sqlDB, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Ping worked ")

	return &store{
		db: sqlDB,
	}, nil
}

func (r *store) Close() {
	r.db.Close()
}

func (r *store) Category() storage.CategoryRepoI {

	if r.category == nil {
		r.category = NewCategoryRepo(r.db)
	}

	return r.category
}

func (r *store) Product() storage.ProductRepoI {
	if r.product == nil {
		r.product = NewProductRepo(r.db)
	}
	return r.product
}