package repository

import (
	"auth-service/internal/model"
	"context"

	"auth-service/internal/db" // Import package db

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository struct holds the MongoDB client
type UserRepository struct {
	Client *mongo.Client
}

// NewUserRepository initializes a new UserRepository
func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		Client: client,
	}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user model.User) error {
	_, err := r.Client.Database(db.DatabaseName).Collection("users").InsertOne(context.TODO(), user)
	return err
}

// FindUserByUsername retrieves a user by their username
func (r *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.Client.Database(db.DatabaseName).Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindUserByChatTelegramID(chatTelegramID int64) (*model.User, error) {
	var user model.User
	err := r.Client.Database(db.DatabaseName).Collection("users").FindOne(
		context.TODO(),
		bson.M{"chat_telegram_id": chatTelegramID},
	).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
