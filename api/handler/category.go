package handler

import (
	"net/http"

	"psql/models"
	"psql/pkg/helper"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "Category to create"
// @Success 200 {object} models.Category "Category created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /category [post]
func (h *handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		h.handlerResponse(c, "shoudBindJSON createCategory", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Category().Create(&createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.create", http.StatusInternalServerError, err.Error())
	}

	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "category GetBydID in create category: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(c, "create category response", http.StatusOK, resp)
}

// GetByIdCategory godoc
// @Summary Get a category by ID
// @Description Get details of a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category "Category details"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal server error"
// @Router /category/{id} [get]
func (h *handler) GetByIdCategory(c *gin.Context) {
	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "invalid id", http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "category GetBydID in get by id category: ", http.StatusInternalServerError, err.Error())
	}

	h.handlerResponse(c, "get by id category response", http.StatusOK, resp)
}

// GetListCategory godoc
// @Summary Get a list of categories
// @Description Get a list of categories with pagination and search options
// @Tags categories
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param search query string false "Search term"
// @Success 200 {array} models.Category "List of categories"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /category [get]
func (h *handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffset(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "GetListCategory offset", http.StatusBadRequest, "invalid offset")
	}

	limit, err := h.getLimit(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "GetListCategory limit", http.StatusBadRequest, "invalid limit")
	}

	resp, err := h.strg.Category().GetList(&models.CategoryGetListReq{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "h.strg.Category().GetList(&models.CategoryGetListReq ", http.StatusInternalServerError, err.Error())
	}

	h.handlerResponse(c, "get list category response", http.StatusOK, resp)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update the details of an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.UpdateCategory true "Category to update"
// @Success 200 {object} models.Category "Category updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /category [put]
func (h *handler) UpdateCategory(c *gin.Context) {
	var updateCategory models.UpdateCategory

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		h.handlerResponse(c, "shoudBindJSON udpate category", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Category().Update(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "strg.category.update", http.StatusInternalServerError, err.Error())
	}

	getCategory, err := h.strg.Category().GetById(&models.CategoryPKey{ID: resp.ID})
	if err != nil {
		h.handlerResponse(c, "strg.category.getbyid: ", http.StatusInternalServerError, err.Error())
	}

	h.handlerResponse(c, "udpate category response", http.StatusOK, getCategory)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {string} string "Category deleted successfully"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal server error"
// @Router /category/{id} [delete]
func (h *handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "invalid id", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Category().Delete(&models.CategoryPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.category.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete category response", http.StatusOK, "Category deleted successfully")
}
