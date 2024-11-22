package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"context"
	"errors"
	"os"

	"auth-service/pkg/jwt"
)

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
	return jwt.GenerateToken(ctx, user.Username, os.Getenv("JWT_SECRET_KEY"))
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

	return jwt.GenerateToken(ctx, user.Username, os.Getenv("JWT_SECRET_KEY"))
}
