package routes

import (
	"productivity-project-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	// Public routes (no authentication required)
	public := router.Group("/users")
	{
		public.POST("/register", controllers)                                   // Sign up(auto-sends verufication code)
		public.POST("/login", controllers.Login)                                // Sign in
		public.POST("/resend-verification", controllers.ResendVerificationCode) // Resend code
		public.POST("/verify-email", controllers.VerifyEmail)                   // Submit verification code
	}
}
