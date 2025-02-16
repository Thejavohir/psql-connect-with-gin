package api

import (
	_ "psql/api/docs"
	"psql/api/handler"
	"psql/config"
	"psql/pkg/logger"
	"psql/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.Logger, cache storage.CacheI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())

	handler := handler.NewHandler(cfg, storage, logger, cache)

	v1 := r.Group("/v1")

	//REGISTER
	r.POST("/register", handler.Register)

	//LOGIN
	r.POST("login", handler.Login)

	//CATEGORY
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category", handler.UpdateCategory)
	r.PATCH("/category/:id", handler.PatchCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	//PRODUCT
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.PatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)
	
	//USER
	v1.Use(handler.AuthMiddleware())

	v1.GET("/user/:id", handler.GetByIdUser)
	v1.GET("/user", handler.GetListUser)
	v1.PUT("/user", handler.UpdateUser)
	v1.PATCH("/user/:id", handler.PatchUser)
	v1.DELETE("/user/:id", handler.DeleteUser)

	//BRANCH
	v1.POST("/branch", handler.CreateBranch)
	v1.GET("/branch/:id", handler.GetByIdBranch)
	v1.GET("/branch", handler.GetListBranch)
	v1.PUT("/branch", handler.UpdateBranch)
	v1.PATCH("/branch/:id", handler.PatchBranch)
	v1.DELETE("/branch/:id", handler.DeleteBranch)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accept-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
