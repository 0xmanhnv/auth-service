package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MongoDBURI       string
	RedisAddress     string
	DatabaseName     string
	TelegramBotToken string
}

func LoadConfig() (*Configuration, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println(os.Getenv("MONGODB_URI"))

	return &Configuration{
		MongoDBURI:       os.Getenv("MONGODB_URI"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		RedisAddress:     os.Getenv("REDIS_ADDRESS"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}, nil
}
