package handler

import (
	"log"
	"net/http"

	"psql/api/models"
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
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /product [post]
func (h *handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	if err := c.ShouldBindJSON(&createProduct); err != nil {
		h.handlerResponse(c, "should bind json in create product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Product().Create(c.Request.Context(), &createProduct)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(c, "strg.product.create", http.StatusInternalServerError, "error creating product")
		return
	}

	resp, err := h.strg.Product().GetById(c.Request.Context(), &models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.product.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product response", http.StatusOK, resp)
}

// GetByIdProduct godoc
// @Summary Get Product by ID
// @Description Get details of a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product "Product details"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal error"
// @Router /product/{id} [get]
func (h *handler) GetByIdProduct(c *gin.Context) {
	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "IsValidUUID", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Product().GetById(c.Request.Context(), &models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.product.GetById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

// GetListProduct godoc
// @Summary Get List of Products
// @Description Get a list of products with pagination and search
// @Tags products
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Success 200 {array} models.Product "List of products"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /product [get]
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

	resp, err := h.strg.Product().GetList(c.Request.Context(), &models.ProductGetListReq{
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

// UpdateProduct godoc
// @Summary Update Product
// @Description Update details of an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.UpdateProduct true "Product to update"
// @Success 200 {object} models.Product "Product updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /product [put]
func (h *handler) UpdateProduct(c *gin.Context) {
	var updateProduct models.UpdateProduct

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		h.handlerResponse(c, "shoudBindJSON udpate Product", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Product().Update(c.Request.Context(), &updateProduct)
	if err != nil {
		h.handlerResponse(c, "strg.Product.update", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.Product.update", http.StatusInternalServerError, "no rows affected")
	}

	getProduct, err := h.strg.Product().GetById(c.Request.Context(), &models.ProductPKey{ID: updateProduct.ID})
	if err != nil {
		h.handlerResponse(c, "strg.Product.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "udpate Product response", http.StatusOK, getProduct)
}


// PatchProduct godoc
// @ID patch_product
// @Summary Patch Product
// @Description Patch details of an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param product body models.PatchRequest true "Product to Patch"
// @Success 200 {object} models.Product "Product Patched successfully"
// @Failure 400 {string} Response{data=string} "Invalid request"
// @Failure 500 {string} Response{data=string} "Internal error"
// @Router /product/{id} [patch]
func (h *handler) PatchProduct(c *gin.Context) {
	var (
		id string = c.Param("id")
		patchProduct models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "isValidUUID", http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := c.ShouldBindJSON(&patchProduct); err != nil {
		h.handlerResponse(c, "shoudBindJSON patch Product", http.StatusBadRequest, err.Error())
		return
	}

	patchProduct.ID = id
	resp, err := h.strg.Product().Patch(c.Request.Context(), &patchProduct)
	if err != nil {
		h.handlerResponse(c, "strg.Product.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.Product.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	getProduct, err := h.strg.Product().GetById(c.Request.Context(), &models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.Product.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "patch Product response", http.StatusOK, getProduct)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 500 {string} string "Internal error"
// @Router /product/{id} [delete]
func (h *handler) DeleteProduct(c *gin.Context) {
	var delProduct models.ProductPKey

	id := c.Param("id")
	delProduct.ID = id

	err := h.strg.Product().Delete(c.Request.Context(), &delProduct)
	if err != nil {
		h.handlerResponse(c, "strg.product.delete", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "delete product response", http.StatusOK, "Deleted successfully")
}
