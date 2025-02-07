package handler

import (
	"net/http"
	"psql/models"

	"github.com/gin-gonic/gin"
)

// func (h *handler) CreateCategory(c *gin.Context) {

// 	var createCategory models.CreateCategory

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		h.handlerResponse(c, "error while reading ReadAll: "+err.Error(), http.StatusBadRequest, nil)
// 	}

// 	err = json.Unmarshal(body, &createCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while unmarshalling in CreateCategory: "+err.Error(), http.StatusInternalServerError, nil)
// 	}

// 	id, err := h.strg.Category().Create(&createCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while creating category in CreateCategory: "+err.Error(), http.StatusInternalServerError, nil)
// 	}

// 	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
// 	if err != nil {
// 		h.handlerResponse(c, "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
// 	}

// 	h.handlerResponse(c, "Success", http.StatusOK, resp)
// }

// func (h *handler) GetByIdCategory(c *gin.Context) {
// 	var id string = r.URL.Query().Get("id")

// 	if !helper.IsValidUUID(id) {
// 		h.handlerResponse(c, "invalid id", http.StatusBadRequest, nil)
// 		return
// 	}
// 	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
// 	if err != nil {
// 		h.handlerResponse(c, "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
// 	}

// 	h.handlerResponse(c, "Success", http.StatusOK, resp)
// }

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

// func (h *handler) UpdateCategory(c *gin.Context) {
// 	var updateCategory models.UpdateCategory

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		h.handlerResponse(c, "error while reading ReadAll in update: "+err.Error(), http.StatusBadRequest, nil)
// 	}

// 	err = json.Unmarshal(body, &updateCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while unmarshalling in UpdateCategory: "+err.Error(), http.StatusInternalServerError, nil)
// 	}

// 	resp, err := h.strg.Category().Update(&updateCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while updating category: "+err.Error(), http.StatusInternalServerError, nil)
// 	}

// 	getCategory, err := h.strg.Category().GetById(&models.CategoryPKey{ID: resp.ID})
// 	if err != nil {
// 		h.handlerResponse(c, "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
// 	}

// 	h.handlerResponse(c, "Success", http.StatusOK, getCategory)
// }

// func (h *handler) DeleteCategory(c *gin.Context) {
// 	var delCategory models.CategoryPKey

// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		h.handlerResponse(c, "error while reading ReadAll in delete: "+err.Error(), http.StatusBadRequest, nil)
// 	}

// 	err = json.Unmarshal(body, &delCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while unmarshalling in DeleteCategory: "+err.Error(), http.StatusInternalServerError, nil)
// 	}

// 	err = h.strg.Category().Delete(&delCategory)
// 	if err != nil {
// 		h.handlerResponse(c, "error while deleting category: ", http.StatusInternalServerError, nil)
// 	}
// 	h.handlerResponse(c, "Success", http.StatusOK, nil)
// }
