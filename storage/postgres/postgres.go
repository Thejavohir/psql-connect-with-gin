package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	"psql/config"
	"psql/storage"
)

type store struct {
	db       *pgxpool.Pool
	category *categoryRepo
	product  *productRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	fmt.Println("Ping worked ")

	return &store{
		db: pgxpool,
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
