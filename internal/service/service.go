package service

import "auth-service/internal/model"

type UserRepository interface {
	CreateUser(user model.User) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserByChatTelegramID(chatTelegramID int64) (*model.User, error)
}
