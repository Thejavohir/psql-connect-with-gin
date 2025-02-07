package api

import (
	"psql/api/handler"
	"psql/config"
	"psql/pkg/logger"
	"psql/storage"

	"github.com/gin-gonic/gin"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.Logger) {

	handler := handler.NewHandler(cfg, storage, logger)

	r.GET("/category", handler.GetListCategory)
}
