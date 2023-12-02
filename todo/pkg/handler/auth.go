package handler

import (
	"net/http"
	models "todo/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {

	var user models.CreateUser

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) signIn(c *gin.Context) {

}
