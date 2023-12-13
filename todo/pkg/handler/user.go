package handler

import (
	"fmt"
	"net/http"
	"strconv"
	models "todo/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c *gin.Context) {

	var user models.CreateUserReq

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.User.CreateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	users, err := h.services.User.GetUserById(&models.IdRequest{ID: id})
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

	resp, err := h.services.GetAllUsers(&models.GetAllUserReq{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUser

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	user.ID = id
	fmt.Println("before", user)

	resp, err := h.services.User.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	fmt.Println("after", resp)
	c.JSON(http.StatusOK, gin.H{"message": "User successfully updated", "id": resp})

}

func (h *Handler) DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	resp, err := h.services.User.DeleteUser(&models.IdRequest{ID: id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted", "id": resp})

}
func (h *Handler) CreateUsers(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	var users []models.CreateUserReq
	fmt.Println("before", users)

	if err := c.BindJSON(&users); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// for i := range users {
	// 	users[i].CreatedBy = userId
	// }

	ids, err := h.services.User.CreateUsers(users)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("after handler", users)

	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}
func (h *Handler) UpdateUsers(c *gin.Context) {

	var users []models.UpdateUser

	if err := c.ShouldBindJSON(&users); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.User.UpdateUsers(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Users successfully updated", "result": result})
}
