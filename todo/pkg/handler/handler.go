package handler

import (
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	user := router.Group("/user")
	{
		user.POST("/", h.CreateUser)
		user.POST("/create/users", h.CreateUsers)
		user.GET("/:id", h.GetUserById)
		user.GET("/", h.GetAllUsers)
		user.PUT("/:id", h.UpdateUser)
		user.PUT("/users", h.UpdateUsers)
		user.DELETE("/:id", h.DeleteUser)
	}

	// api := router.Group("/api")
	// {
	// 	lists := api.Group("/lists")
	// 	{
	// 		lists.POST("/", h.createList)
	// 		lists.GET("/", h.getAllLists)
	// 		lists.GET("/:id", h.getListById)
	// 		lists.PUT("/:id", h.updateList)
	// 		lists.DELETE("/:id", h.deleteList)

	// 		items := lists.Group(":id/items")
	// 		{
	// 			items.POST("/", h.createItem)
	// 			items.GET("/", h.getAllItems)
	// 		}
	// 	}


	// }

	return router
}
