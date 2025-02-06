package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"psql/config"
	"psql/storage"
)

type handler struct {
	cfg *config.Config
	strg storage.StorageI
}

type response struct {
	Status int `json:"status"`
	Description string `json:"description"`
	Data interface{}
}

func NewHandler(cfg *config.Config, storage storage.StorageI) *handler {
	return &handler{
		cfg: cfg,
		strg: storage,
	}
}

func (h *handler) handlerResponse(w http.ResponseWriter, message string, code int, data interface{}) {
	resp := response{
		Status: code,
		Description: message,
		Data: data,
	}

	log.Printf("%+v", resp)

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)

	}

	fmt.Println()

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}