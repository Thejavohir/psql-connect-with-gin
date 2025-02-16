package handler

import (
	"net/http"

	"psql/api/models"
	"psql/pkg/helper"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// GetByIdUser godoc
// @Summary Get User by ID
// @Description Get details of a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User "User details"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal error"
// @Router /v1/user/{id} [get]
func (h *handler) GetByIdUser(c *gin.Context) {
	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "IsValidUUID", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.User().GetById(c.Request.Context(), &models.UserPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.user.GetById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// GetListUser godoc
// @Summary Get List of Users
// @Description Get a list of users with pagination and search
// @Tags users
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Param search_barcode query int false "Search by Barcode"
// @Success 200 {array} models.User "List of users"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /v1/user [get]
func (h *handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffset(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "GetListUser offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimit(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "GetListUser limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListReq{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "h.strg.User().GetList(&models.UserGetListReq ", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "getById response", http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update details of an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UpdateUser true "User to update"
// @Success 200 {object} models.User "User updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /v1/user [put]
func (h *handler) UpdateUser(c *gin.Context) {
	var updateUser models.UpdateUser

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		h.handlerResponse(c, "shoudBindJSON udpate User", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().Update(c.Request.Context(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "strg.User.update", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.User.update", http.StatusInternalServerError, "no rows affected")
	}

	getUser, err := h.strg.User().GetById(c.Request.Context(), &models.UserPKey{ID: updateUser.ID})
	if err != nil {
		h.handlerResponse(c, "strg.User.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "udpate User response", http.StatusOK, getUser)
}


// PatchUser godoc
// @ID patch_user
// @Summary Patch User
// @Description Patch details of an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.PatchRequest true "User to Patch"
// @Success 200 {object} models.User "User Patched successfully"
// @Failure 400 {string} Response{data=string} "Invalid request"
// @Failure 500 {string} Response{data=string} "Internal error"
// @Router /v1/user/{id} [patch]
func (h *handler) PatchUser(c *gin.Context) {
	var (
		id string = c.Param("id")
		patchUser models.PatchRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "isValidUUID", http.StatusBadRequest, "invalid uuid")
		return
	}

	if err := c.ShouldBindJSON(&patchUser); err != nil {
		h.handlerResponse(c, "shoudBindJSON patch User", http.StatusBadRequest, err.Error())
		return
	}

	patchUser.ID = id
	resp, err := h.strg.User().Patch(c.Request.Context(), &patchUser)
	if err != nil {
		h.handlerResponse(c, "strg.User.patch", http.StatusInternalServerError, err.Error())
		return
	}

	if resp <= 0 {
		h.handlerResponse(c, "strg.User.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	getUser, err := h.strg.User().GetById(c.Request.Context(), &models.UserPKey{ID: id})
	if err != nil {
		h.handlerResponse(c, "strg.User.getbyid: ", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "patch User response", http.StatusOK, getUser)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 500 {string} string "Internal error"
// @Router /v1/user/{id} [delete]
func (h *handler) DeleteUser(c *gin.Context) {
	var delUser models.UserPKey

	id := c.Param("id")
	delUser.ID = id

	err := h.strg.User().Delete(c.Request.Context(), &delUser)
	if err != nil {
		h.handlerResponse(c, "strg.user.delete", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "delete user response", http.StatusOK, "Deleted successfully")
}
