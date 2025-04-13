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
		noteController controllers.NoteController,
		store *sessions.CookieStore,
	) {
	env := os.Getenv("ENV")
	var frontendURL string

	switch env {
	case "production":
		frontendURL = os.Getenv("PROD_FRONTEND_URL")
	default:
		frontendURL = os.Getenv("DEV_FRONTEND_URL")
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// TODO: make this only applying on extraneous queries
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  2000,
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
		public.GET("/current-user", authController.GetCurrentUser)
		public.POST("/logout", authController.Logout)
	}

	protected := router.Group("/notes")
	protected.Use(rateLimiter)
	protected.Use(middlewares.AuthMiddleware(store))
	{
		protected.GET("", noteController.GetAllNotes)
		protected.POST("", noteController.CreateNote)
		protected.PUT("/:id", noteController.UpdateNote)
	}
}
