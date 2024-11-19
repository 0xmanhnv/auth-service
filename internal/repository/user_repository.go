package repository

import (
	"auth-service/internal/config"
	"auth-service/internal/model"
	"context"

	"auth-service/internal/db" // Import package db

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository struct holds the MongoDB client
type UserRepository struct {
	client *mongo.Client
}

// NewUserRepository initializes a new UserRepository
func NewUserRepository(cfg *config.Configuration) (*UserRepository, error) {
	client := db.Client // Use the client from the db package
	return &UserRepository{client: client}, nil
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user model.User) error {
	_, err := r.client.Database(db.DatabaseName).Collection("users").InsertOne(context.TODO(), user)
	return err
}

// FindUserByUsername retrieves a user by their username
func (r *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.client.Database(db.DatabaseName).Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
