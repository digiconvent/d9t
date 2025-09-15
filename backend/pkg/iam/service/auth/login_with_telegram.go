package iam_auth_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IamAuthService) LoginTelegramUser(body, botToken string) (*uuid.UUID, *core.Status) {
	telegramId, status := s.repository.User.GetTelegramID(body, botToken)
	if status.Err() {
		return nil, &status
	}

	userId, status := s.repository.User.GetUserByTelegramID(telegramId)
	if status.Err() {
		return nil, &status
	}

	return userId, &status
}
