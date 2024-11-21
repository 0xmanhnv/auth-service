package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	Bot      *tgbotapi.BotAPI
	userRepo *repository.UserRepository
}

func NewTelegramService(token string, userRepo *repository.UserRepository) (*TelegramService, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &TelegramService{
		Bot:      bot,
		userRepo: userRepo,
	}, nil
}

func (s *TelegramService) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-message updates
			continue
		}

		user := model.User{
			Username:       update.Message.From.UserName,
			ChatTelegramID: update.Message.Chat.ID,
		}

		// Lưu thông tin người dùng
		err := s.userRepo.CreateUser(user)
		if err != nil {
			log.Printf("Error saving user: %v", err)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bạn đã đăng nhập thành công!")
		s.Bot.Send(msg)
	}
}
