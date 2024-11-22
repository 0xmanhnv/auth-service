package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username"`
	Password       string             `bson:"password"`
	ChatTelegramID int64              `bson:"chat_telegram_id"`
	FirstName      string             `bson:"first_name"`
	PhotoUrl       string             `bson:"photo_url"`
}

type LoginRequest struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
