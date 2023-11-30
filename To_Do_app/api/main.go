package api

import (
	_ "user/api/docs"
	"user/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	//Users

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.GET("/user", h.GetAllUser)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user", h.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
