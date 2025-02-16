package handler

import (
	"net/http"
	"psql/pkg/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")

		result, err := helper.ParseClaims(value, h.cfg.PrivateKey)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Set("Auth", result)

		c.Next()
	}
}
