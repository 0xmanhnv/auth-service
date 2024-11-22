package server

import (
	"auth-service/docs"
	"auth-service/internal/config"
	"auth-service/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	port int
	db   *mongo.Client
}

func NewServer() *http.Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	NewServer := &Server{
		port: port,
		db:   db.InitMongoDB(cfg.MongoDBURI), // InitMongoDB returns a MongoDB client
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.GinServer(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) GinServer() http.Handler {

	// Set gin mode
	if os.Getenv("APP_MODE") == "release" || os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init gin default
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cfg.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Description = "Auth Service API"
	docs.SwaggerInfo.Title = "Auth Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", s.port)

	// Register routes
	s.RegisterRoutes(r)

	// swagger
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(1),
			ginSwagger.DeepLinking(true),
			ginSwagger.PersistAuthorization(true),
		),
	)

	return r
}
