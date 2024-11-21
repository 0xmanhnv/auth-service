package server

import (
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

	// Init gin default
	ginServer := NewServer.GinServer(cfg)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(ginServer),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) GinServer(cfg *config.Configuration) *gin.Engine {
	// Init gin default
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	return r
}
