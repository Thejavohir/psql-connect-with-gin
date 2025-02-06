package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"psql/models"
	"psql/pkg/helper"
	"strconv"
)


func (h *handler) Category(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		h.CreateCategory(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET_LIST" {
			h.GetListCategory(w, r)
		}else if method == "GET" {
			h.GetByIdCategory(w, r)
		}
	case "PUT":
		h.UpdateCategory(w, r)
	case "DELETE":
		h.DeleteCategory(w, r)
	}

}

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var createCategory models.CreateCategory

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while reading ReadAll: " + err.Error(), http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, &createCategory)
	if err != nil {
		h.handlerResponse(w, "error while unmarshalling in CreateCategory: " + err.Error(), http.StatusInternalServerError, nil)
	}

	id, err := h.strg.Category().Create(&createCategory)
	if err != nil {
		h.handlerResponse(w, "error while creating category in CreateCategory: " + err.Error(), http.StatusInternalServerError, nil)
	}

	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
	if err != nil {
		h.handlerResponse(w,  "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) GetByIdCategory(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "invalid id", http.StatusBadRequest, nil)
		return
	}
	resp, err := h.strg.Category().GetById(&models.CategoryPKey{ID: id})
	if err != nil {
		h.handlerResponse(w,  "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) GetListCategory(w http.ResponseWriter, r *http.Request) {

	var (
		offsetStr = r.URL.Query().Get("offset")
		limitStr = r.URL.Query().Get("limit")

		search = r.URL.Query().Get("search")
	)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		h.handlerResponse(w, "error offset: " + err.Error(), http.StatusBadRequest, nil)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.handlerResponse(w, "error limit: " + err.Error(), http.StatusBadRequest, nil)
	}

	resp, err := h.strg.Category().GetList(&models.CategoryGetListReq{
		Offset: offset,
		Limit: limit,
		Search: search,
	})
	if err != nil {
		h.handlerResponse(w,  "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateCategory models.UpdateCategory

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while reading ReadAll in update: " + err.Error(), http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, &updateCategory)
	if err != nil {
		h.handlerResponse(w, "error while unmarshalling in UpdateCategory: " + err.Error(), http.StatusInternalServerError, nil)
	}

	resp, err := h.strg.Category().Update(&updateCategory)
	if err != nil {
		h.handlerResponse(w, "error while updating category: " + err.Error(), http.StatusInternalServerError, nil)
	}

	getCategory, err := h.strg.Category().GetById(&models.CategoryPKey{ID: resp.ID})
	if err != nil {
		h.handlerResponse(w,  "error while getting category in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, getCategory)
}

func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var delCategory models.CategoryPKey
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while reading ReadAll in delete: " + err.Error(), http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, &delCategory)
	if err != nil {
		h.handlerResponse(w, "error while unmarshalling in DeleteCategory: " + err.Error(), http.StatusInternalServerError, nil)
	}

	err = h.strg.Category().Delete(&delCategory)
	if err != nil {
		h.handlerResponse(w,  "error while deleting category: ", http.StatusInternalServerError, nil)
	}
	h.handlerResponse(w, "Success", http.StatusOK, nil)
}