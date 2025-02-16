package handler

import (
	"fmt"
	"log"
	"net/http"

	"psql/api/models"
	"psql/config"
	"psql/pkg/helper"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @Summary Create Branch
// @Description Create a new branch with provided details
// @Tags branches
// @Accept json
// @Produce json
// @Param branch body models.CreateBranch true "Branch to create"
// @Success 200 {object} models.Branch "Branch created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /v1/branch [post]
func (h *handler) CreateBranch(c *gin.Context) {

	var createBranch models.CreateBranch

	if err := c.ShouldBindJSON(&createBranch); err != nil {
		h.handlerResponse(c, "should bind json in create branch", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Branch().Create(c.Request.Context(), &createBranch)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(c, "strg.branch.create", http.StatusInternalServerError, "error creating branch")
		return
	}

	// resp, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPKey{ID: id})
	// if err != nil {
	// 	h.handlerResponse(c, "strg.branch.getbyid: ", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	h.handlerResponse(c, "create branch response", http.StatusOK, id)
}

// GetByIdBranch godoc
// @Summary Get Branch by ID
// @Description Get details of a branch by its ID
// @Tags branches
// @Accept json
// @Produce json
// @Param id path string true "Branch ID"
// @Success 200 {object} models.Branch "Branch details"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal error"
// @Router /v1/branch/{id} [get]
func (h *handler) GetByIdBranch(c *gin.Context) {
	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "IsValidUUID", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.branch.GetById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// GetListBranch godoc
// @Summary Get List of Branches
// @Description Get a list of branches with pagination and search
// @Tags branches
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Success 200 {array} models.Branch "List of branches"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /v1/branch [get]
func (h *handler) GetListBranch(c *gin.Context) {

	info, exitst := c.Get("Auth")
	if !exitst {
		h.handlerResponse(c, "get('Auth')", http.StatusBadRequest, "invalid info")
		return
	}

	user := info.(helper.TokenInfo)
	fmt.Println(user)
	if user.ClientType == config.SuperAdmin {
		user.UserID = ""
	}

	offset, err := h.getOffset(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "GetListBranch offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimit(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "GetListBranch limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Branch().GetList(c.Request.Context(), &models.BranchGetListReq{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "h.strg.Branch().GetList(&models.BranchGetListReq ", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "get list of branches response", http.StatusOK, resp)
}

// UpdateBranch godoc
// @Summary Update Branch
// @Description Update details of an existing branch
// @Tags branches
// @Accept json
// @Produce json
// @Param branch body models.UpdateBranch true "Branch to update"
// @Success 200 {object} models.Branch "Branch updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /v1/branch [put]
func (h *handler) UpdateBranch(c *gin.Context) {
	var updateBranch models.UpdateBranch

	if err := c.ShouldBindJSON(&updateBranch); err != nil {
		h.handlerResponse(c, "shoudBindJSON udpate Branch", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Branch().Update(c.Request.Context(), &updateBranch)
	if err != nil {
		h.handlerResponse(c, "strg.Branch.update", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.Branch.update", http.StatusInternalServerError, "no rows affected")
	}

	getBranch, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPKey{ID: updateBranch.ID})
	if err != nil {
		h.handlerResponse(c, "strg.Branch.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "udpate Branch response", http.StatusOK, getBranch)
}

// PatchBranch godoc
// @ID patch_branch
// @Summary Patch Branch
// @Description Patch details of an existing branch
// @Tags branches
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param branch body models.PatchRequest true "Branch to Patch"
// @Success 200 {object} models.Branch "Branch Patched successfully"
// @Failure 400 {string} Response{data=string} "Invalid request"
// @Failure 500 {string} Response{data=string} "Internal error"
// @Router /v1/branch/{id} [patch]
func (h *handler) PatchBranch(c *gin.Context) {
	var (
		id          string = c.Param("id")
		patchBranch models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "isValidUUID", http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := c.ShouldBindJSON(&patchBranch); err != nil {
		h.handlerResponse(c, "shoudBindJSON patch Branch", http.StatusBadRequest, err.Error())
		return
	}

	patchBranch.ID = id
	resp, err := h.strg.Branch().Patch(c.Request.Context(), &patchBranch)
	if err != nil {
		h.handlerResponse(c, "strg.Branch.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.Branch.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	getBranch, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.Branch.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "patch Branch response", http.StatusOK, getBranch)
}

// DeleteBranch godoc
// @Summary Delete Branch
// @Description Delete a branch by its ID
// @Tags branches
// @Accept json
// @Produce json
// @Param id path string true "Branch ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 500 {string} string "Internal error"
// @Router /v1/branch/{id} [delete]
func (h *handler) DeleteBranch(c *gin.Context) {
	var delBranch models.BranchPKey

	id := c.Param("id")
	delBranch.ID = id

	err := h.strg.Branch().Delete(c.Request.Context(), &delBranch)
	if err != nil {
		h.handlerResponse(c, "strg.branch.delete", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "delete branch response", http.StatusOK, "Deleted successfully")
}
