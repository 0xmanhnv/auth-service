package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI       string
	RedisAddress     string
	DatabaseName     string
	TelegramBotToken string
	JWTSecretKey     string
	AllowOrigins     []string
}

var Cfg *Config

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Cfg = &Config{
		MongoDBURI:       os.Getenv("MONGODB_URI"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		RedisAddress:     os.Getenv("REDIS_ADDRESS"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		JWTSecretKey:     os.Getenv("JWT_SECRET_KEY"),
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080", "*"},
	}
	fmt.Println("Config loaded!!!")

	return Cfg, nil
}
