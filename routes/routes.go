package routes

import (
	"github.com/gin-gonic/gin"
	"messanger-backend/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/profile", controllers.Profile)
	}
}