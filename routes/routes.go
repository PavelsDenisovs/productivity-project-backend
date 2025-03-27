package routes

import (
	"productivity-project-backend/controllers"
	"github.com/gin-gonic/gin"

	"os"
	"time"
	"github.com/ulule/limiter/v3"
	limiterGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	limiterStore "github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/gin-contrib/cors"
)

func RegisterRoutes(router *gin.Engine, authMiddleware gin.HandlerFunc) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
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

	public := router.Group("/users")
	public.Use(rateLimiter)
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
		public.POST("/resend-verification", controllers.ResendVerificationCode)
		public.POST("/verify-email", controllers.VerifyEmail)
	}

	auth := router.Group("/")
	auth.Use(authMiddleware)
	{
		auth.POST("/logout", controllers.Logout)
	}
}
