package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	var user data.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(user)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {
	var userSing data.SingInInput

	if err := c.BindJSON(&userSing); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(userSing.Login, userSing.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := h.services.Authorization.GetUserIdByLoginAndPass(userSing.Login, userSing.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.AuthEvent.AddAuthenticationEvent(userId, c.Request.Method, c.Request.RequestURI); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
