package handler

import (
	"log"
	"net/http"

	"psql/models"
	"psql/pkg/helper"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create Product
// @Description Create a new product with provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.CreateProduct true "Product to create"
// @Success 200 {object} models.Product "Product created successfully"
// @Failure 400 {string} "Invalid request"
// @Failure 500 {string} "Internal error"
// @Router /product [post]
func (h *handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	if err := c.ShouldBindJSON(&createProduct); err != nil {
		h.handlerResponse(c, "should bind json in create product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Product().Create(&createProduct)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(c, "strg.product.create", http.StatusInternalServerError, "error creating product")
		return
	}

	resp, err := h.strg.Product().GetById(&models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.product.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product response", http.StatusOK, resp)
}

func (h *handler) GetByIdProduct(c *gin.Context) {
	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "IsValidUUID", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Product().GetById(&models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.product.GetById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

func (h *handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffset(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "GetListProduct offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimit(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "GetListProduct limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Product().GetList(&models.ProductGetListReq{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "h.strg.Product().GetList(&models.ProductGetListReq ", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

func (h *handler) UpdateProduct(c *gin.Context) {
	var updateProduct models.UpdateProduct

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		h.handlerResponse(c, "shoudBindJSON udpate Product", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Product().Update(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "strg.Product.update", http.StatusInternalServerError, err.Error())
		return
	}

	getProduct, err := h.strg.Product().GetById(&models.ProductPKey{ID: resp.ID})
	if err != nil {
		h.handlerResponse(c, "strg.Product.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "udpate Product response", http.StatusOK, getProduct)
}

func (h *handler) DeleteProduct(c *gin.Context) {
	var delProduct models.ProductPKey

	id := c.Param("id")
	delProduct.ID = id

	err := h.strg.Product().Delete(&delProduct)
	if err != nil {
		h.handlerResponse(c, "strg.product.delete", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "delete product response", http.StatusOK, "Deleted successfully")
}
