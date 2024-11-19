package service

import (
	"auth-service/internal/model"
	"auth-service/internal/redis"
	"auth-service/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	redisClient *redis.RedisClient
}

func NewAuthService(userRepo *repository.UserRepository, redisClient *redis.RedisClient) *AuthService {
	return &AuthService{userRepo: userRepo, redisClient: redisClient}
}

func (s *AuthService) Register(user model.User) error {
	existingUser, _ := s.userRepo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}
	return s.GenerateToken(user.Username)
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key")) // Thay "your_secret_key" bằng khóa bí mật của bạn
}
