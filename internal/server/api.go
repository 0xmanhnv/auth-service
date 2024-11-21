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
	handler := handler.NewAuthHandler(authSvc)

	// Define API routes
	v1Routes := r.Group("/api/v1")
	{

		authRoutes := v1Routes.Group("/auth")
		{
			authRoutes.GET("/telegram", handler.LoginWithTelegramHandler)

			authRoutes.POST("/register", handler.RegisterHandler)
			authRoutes.POST("/login", handler.LoginHandler)
			authRoutes.GET("/protected", middleware.ValidateToken(), handler.ProtectedHandler)
		}

	}
	return r
}
