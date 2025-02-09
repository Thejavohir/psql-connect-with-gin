package api

import (
	"psql/api/handler"
	"psql/config"
	"psql/pkg/logger"
	"psql/storage"
	_ "psql/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.Logger) {

	handler := handler.NewHandler(cfg, storage, logger)

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category", handler.UpdateCategory)
	r.PATCH("/category/:id", handler.PatchCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.PatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
