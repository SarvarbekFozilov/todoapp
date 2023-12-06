package handler

import (
	"net/http"
	"strconv"
	models "todo/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c *gin.Context) {

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
func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	users, err := h.services.Authorization.GetUserById(&models.IdRequest{Id: id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
func (h *Handler) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logrus.Fatalf("Error getting page: %s", err.Error())
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		logrus.Fatalf("Error getting limit: %s", err.Error())
		c.JSON(http.StatusBadRequest, "invalid limit param")
		return
	}

	resp, err := h.services.GetAllUsers(&models.GetAllUserRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user.ID = c.Param("id")
	resp, err := h.services.Authorization.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully updated", "id": resp})

}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.services.Authorization.DeleteUser(&models.IdRequest{Id: id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted", "id": resp})

}
func (h *Handler) CreateUsers(c *gin.Context) {
	var users []models.CreateUser

	if err := c.BindJSON(&users); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := h.services.Authorization.CreateUsers(users)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}
func (h *Handler) UpdateUsers(c *gin.Context) {
	var users []models.User

	if err := c.BindJSON(&users); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := h.services.Authorization.UpdateUsers(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Users successfully updated", "ids": ids})
}
