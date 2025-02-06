package api

import (
	"net/http"
	"psql/api/handler"
	"psql/config"
	"psql/storage"
)

func NewApi(cfg *config.Config, storage storage.StorageI) {

	handler := handler.NewHandler(cfg, storage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/product", handler.Product)
}