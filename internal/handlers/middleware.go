package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authoriationHeader = "Authorization"
)

func (h *Handler) AdminIdentity(c *gin.Context) {
	header := c.GetHeader(authoriationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty request header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid header size")
	}

	//parse token
	id, err := h.services.Authorization.Parsetoken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("admin_id", id)
}
