package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"psql/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var productTestRepo *productRepo

func TestMain(m *testing.M) {

	cfg := config.Load()

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		panic(err)
	}

	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		panic(err)
	}

	productTestRepo = NewProductRepo(pgxpool)

	os.Exit(m.Run())
}
