// exempt from testing
package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IamUserService) ConnectTelegramUser(initData, botToken string, userId *uuid.UUID) *core.Status {
	telegramId, status := s.repository.User.GetTelegramID(initData, botToken)
	if status.Err() {
		return &status
	}

	status = s.repository.User.RegisterTelegramUser(*telegramId, userId)
	if status.Err() {
		return &status
	}

	return &status
}
