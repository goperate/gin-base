package routes

import (
	"app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
}
