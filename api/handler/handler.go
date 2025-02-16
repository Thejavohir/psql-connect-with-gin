package handler

import (
	"net/http"
	"psql/config"
	"psql/pkg/logger"
	"psql/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg    *config.Config
	strg   storage.StorageI
	logger logger.Logger
	cache  storage.CacheI
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, storage storage.StorageI, logger logger.Logger, cache storage.CacheI) *handler {
	return &handler{
		cfg:    cfg,
		strg:   storage,
		logger: logger,
		cache: cache,
	}
}

func (h *handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	resp := Response{
		Status:      code,
		Description: http.StatusText(code),
		Data:        message,
	}

	switch {
	case code < 300:
		h.logger.Info(path, logger.Any("info", resp))
	case code >= 400:
		h.logger.Error(path, logger.Any("error", resp))
	}

	c.JSON(code, resp)
}

func (h *handler) getOffset(offset string) (int, error) {

	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *handler) getLimit(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}
	return strconv.Atoi(limit)
}
