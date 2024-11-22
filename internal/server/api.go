package server

import (
	"auth-service/internal/db"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes
func (s *Server) RegisterRoutes(r *gin.Engine) http.Handler {
	// Define Auth service
	userRepo := repository.NewUserRepository(db.Client)
	authSvc := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	// Define API routes
	v1Routes := r.Group("/api/v1")
	{

		authRoutes := v1Routes.Group("/auth")
		{
			// @Summary Login with Telegram
			// @Description Login using Telegram authentication
			// @Tags auth
			// @Accept json
			// @Produce json
			// @Success 200 {object} map[string]interface{}
			// @Router /auth/telegram [get]
			authRoutes.GET("/telegram", authHandler.LoginWithTelegramHandler)

			// @Summary Register a new user
			// @Description Register a new user in the system
			// @Tags auth
			// @Accept json
			// @Produce json
			// @Param user body handler.RegisterRequest true "User registration details"
			// @Success 201 {object} map[string]interface{}
			// @Router /auth/register [post]
			authRoutes.POST("/register", authHandler.RegisterHandler)

			// @Summary User login
			// @Description Login an existing user
			// @Tags auth
			// @Accept json
			// @Produce json
			// @Param login body model.LoginRequest true "User login details"
			// @Success 200 {object} map[string]interface{}
			// @Router /auth/login [post]
			authRoutes.POST("/login", authHandler.LoginHandler)

			// @Summary Get protected resource
			// @Description Access a protected resource
			// @Tags auth
			// @Accept json
			// @Produce json
			// @Success 200 {object} map[string]interface{}
			// @Router /auth/protected [get]
			// @Security ApiKeyAuth
			// @securityDefinitions.bearer BearerAuth
			// @type apiKey
			// @name Authorization
			// @in header
			authRoutes.GET("/protected", middleware.JWTMiddleware(), authHandler.ProtectedHandler)
		}

	}
	return r
}
