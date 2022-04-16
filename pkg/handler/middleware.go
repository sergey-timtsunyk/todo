package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCxt             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerPart := strings.Split(header, " ")
	if len(headerPart) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParserToken(headerPart[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCxt, userId)
}

func getUserIdFormContext(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCxt)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(uint)
	if !ok {
		return 0, errors.New("user id is invalid type")
	}

	return idInt, nil
}
