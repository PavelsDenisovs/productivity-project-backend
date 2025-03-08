package routes

import (
	"github.com/gin-gonic/gin"
	"messenger-backend/controllers"
)

func RegisterRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	// Public routes (no authentication required)
	public := router.Group("/users")
	{
		public.POST("/register", controllers.Register) // Sign up(auto-sends verufication code)
		public.POST("/login", controllers.Login) // Sign in
		public.POST("/resend-verification", controllers.ResendVerificationCode) // Resend code
		public.POST("/verify-email", controllers.VerifyEmail) // Submit verification code
	}

	// Private routes (require authentication)
	private := router.Group("/users")
	private.Use(authMiddleware)
	{
		private.GET("/me", controllers.GetCurrentUserProfile)
		private.PUT("/me", controllers.UpdateProfile)
		private.GET("", controllers.SearchUsers)
	}
}