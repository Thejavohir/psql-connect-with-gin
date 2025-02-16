package handler

import (
	"log"
	"net/http"
	"psql/api/models"
	"psql/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register
// @Description Register a new user with provided details
// @Tags register
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User to register"
// @Success 200 {object} models.User "Registered successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /register [post]
func (h *handler) Register(c *gin.Context) {

	var registerUser models.CreateUser

	if err := c.ShouldBindJSON(&registerUser); err != nil {
		h.handlerResponse(c, "should bind json in create user", http.StatusBadRequest, err.Error())
		return
	}
	
	resp, err := h.strg.User().GetById(c.Request.Context(), &models.UserPKey{Username: registerUser.Username})
	if err != nil {
		h.handlerResponse(c, "strg.user.getbyid", http.StatusInternalServerError, err.Error())
		return
	}
	
	if resp != nil && resp.Username != "" {
		h.handlerResponse(c, "strg.user.getbyid: ", http.StatusBadRequest, "User already exists")
		return
	}

	id, err := h.strg.User().Create(c.Request.Context(), &registerUser)
	if err != nil {
		log.Printf("%+v: ", err)
		h.handlerResponse(c, "strg.user.create", http.StatusInternalServerError, "error creating user")
		return
	}

	h.handlerResponse(c, "create user response", http.StatusOK, id)
}


// Login godoc
// @Summary Login
// @Description Login
// @Tags login
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "Login"
// @Success 200 {string} string "Login successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal error"
// @Router /login [post]
func (h *handler) Login(c *gin.Context) {
	var loginUser models.CreateUser

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		h.handlerResponse(c, "Bind JSON to loginUser", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().GetByUsername(c.Request.Context(), &models.UserPKey{Username: loginUser.Username})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "strg.user.GetByUsername", http.StatusBadRequest, "incorrect username")
			return
		}else {
			h.handlerResponse(c, "strg.user.GetByUsername", http.StatusInternalServerError, err.Error())
			return
		}
	}

	if resp.Password != loginUser.Password {	
		h.handlerResponse(c, "password not equals password", http.StatusBadRequest, "incorrect password")
		return
	}

	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": resp.ID,
	}, time.Duration(time.Now().Year()), h.cfg.PrivateKey)

	if err != nil {
		h.handlerResponse(c, "generateJWT", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "token response", http.StatusOK, token)
}