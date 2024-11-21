package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserRepository interface {
	CreateUser(user model.User) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserByChatTelegramID(chatTelegramID int64) (*model.User, error)
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, user model.User) error {
	existingUser, _ := s.userRepo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}
	return s.GenerateToken(ctx, user.Username)
}

func (s *AuthService) LoginWithTelegram(ctx context.Context, u model.User) (string, error) {
	user, err := s.userRepo.FindUserByChatTelegramID(u.ChatTelegramID)
	if err != nil {
		// create user
		if err := s.userRepo.CreateUser(u); err != nil {
			return "", err
		}
		user, err = s.userRepo.FindUserByChatTelegramID(u.ChatTelegramID)
		if err != nil {
			return "", err
		}
	}
	return s.GenerateToken(ctx, user.Username)
}

func (s *AuthService) GenerateToken(ctx context.Context, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key")) // Thay "your_secret_key" bằng khóa bí mật của bạn
}
