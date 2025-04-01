package routes

import (
	"productivity-project-backend/controllers"
	"productivity-project-backend/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/ulule/limiter/v3"
	limiterGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	limiterStore "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RegisterRoutes(
		router *gin.Engine, 
		authController controllers.AuthController, 
		verificationController controllers.VerificationController,
		store *sessions.CookieStore,
	) {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  5,
	}	

	// Create middleware
	rateLimiter := limiterGin.NewMiddleware(
		limiter.New(
				limiterStore.NewStore(),
				rate,
				limiter.WithTrustForwardHeader(true),
		),
	)

	public := router.Group("/auth")
	public.Use(rateLimiter)
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.POST("/resend-verification", verificationController.ResendVerification)
		public.POST("/verify-email", verificationController.VerifyEmail)
	}
	// TODO: implement /logout route
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware(store))
	{
		auth.POST("/logout", controllers.Logout)
	}
}
