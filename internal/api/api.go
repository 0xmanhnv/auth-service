package api

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
	"auth-service/internal/redis"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

// Setup api
func SetupApi(r *gin.Engine) {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Khởi tạo kết nối đến MongoDB
	db.InitMongoDB(cfg.MongoDBURI)

	// Khởi tạo Redis client
	redisClient := redis.NewRedisClient(cfg.RedisAddress)

	// Khởi tạo User Repository
	userRepo, err := repository.NewUserRepository(cfg)
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}

	// Khởi tạo Auth Service
	authService := service.NewAuthService(userRepo, redisClient)

	// Khởi tạo Gin router
	router := gin.Default()

	// Thiết lập các route
	SetupRoutes(r, authService)

	log.Println("Auth service is running on :8080")
	log.Fatal(router.Run(":8080"))
}

// SetupRoutes thiết lập các route cho API
func SetupRoutes(r *gin.Engine, authService *service.AuthService) {
	handler := handler.NewAuthHandler(authService)
	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)
	r.GET("/protected", middleware.ValidateToken(), handler.ProtectedHandler)
}
