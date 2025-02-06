package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"psql/models"
	"psql/pkg/helper"
)

func (h *handler) Product(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		h.CreateProduct(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)
		if method == "GET_LIST" {
			h.GetListProduct(w, r)
		} else if method == "GET" {
			h.GetByIdProduct(w, r)
		}
	case "PUT":
		h.UpdateProduct(w, r)
	case "DELETE":
		h.DeleteProduct(w, r)
	}
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error on ReadAll in CreateProduct", http.StatusBadRequest, nil)
		return
	}

	var createProduct models.CreateProduct
	err = json.Unmarshal(body, &createProduct)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(w, "error while unmarshalling in CreateProduct", http.StatusInternalServerError, nil)
		return
	}

	id, err := h.strg.Product().Create(&createProduct)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(w, "error while creating product", http.StatusInternalServerError, nil)
		return
	}

	resp, err := h.strg.Product().GetById(&models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(w, "error while getting Product in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) GetByIdProduct(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(w, "invalid id", http.StatusBadRequest, nil)
		return
	}
	resp, err := h.strg.Product().GetById(&models.ProductPKey{ID: id})
	if err != nil {
		h.handlerResponse(w, "error while getting Product in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) GetListProduct(w http.ResponseWriter, r *http.Request) {

	var (
		offsetStr = r.URL.Query().Get("offset")
		limitStr  = r.URL.Query().Get("limit")

		search = r.URL.Query().Get("search")
	)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		h.handlerResponse(w, "error offset: "+err.Error(), http.StatusBadRequest, nil)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.handlerResponse(w, "error limit: "+err.Error(), http.StatusBadRequest, nil)
	}

	resp, err := h.strg.Product().GetList(&models.ProductGetListReq{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		h.handlerResponse(w, "error while getting Product in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, resp)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct models.UpdateProduct

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while reading ReadAll in update: "+err.Error(), http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, &updateProduct)
	if err != nil {
		h.handlerResponse(w, "error while unmarshalling in UpdateProduct: "+err.Error(), http.StatusInternalServerError, nil)
	}

	resp, err := h.strg.Product().Update(&updateProduct)
	if err != nil {
		h.handlerResponse(w, "error while updating Product: "+err.Error(), http.StatusInternalServerError, nil)
	}

	getProduct, err := h.strg.Product().GetById(&models.ProductPKey{ID: resp.ID})
	if err != nil {
		h.handlerResponse(w, "error while getting Product in GetBydID: ", http.StatusInternalServerError, nil)
	}

	h.handlerResponse(w, "Success", http.StatusOK, getProduct)
}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var delProduct models.ProductPKey

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handlerResponse(w, "error while reading ReadAll in delete: "+err.Error(), http.StatusBadRequest, nil)
	}

	err = json.Unmarshal(body, &delProduct)
	if err != nil {
		h.handlerResponse(w, "error while unmarshalling in DeleteProduct: "+err.Error(), http.StatusInternalServerError, nil)
	}

	err = h.strg.Product().Delete(&delProduct)
	if err != nil {
		h.handlerResponse(w, "error while deleting Product: ", http.StatusInternalServerError, nil)
	}
	h.handlerResponse(w, "Success", http.StatusOK, nil)
}
