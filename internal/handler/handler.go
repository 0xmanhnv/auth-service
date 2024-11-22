package handler

import (
	"auth-service/internal/model"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, username, password string) (string, error)
	LoginWithTelegram(ctx context.Context, u model.User) (string, error)
}
