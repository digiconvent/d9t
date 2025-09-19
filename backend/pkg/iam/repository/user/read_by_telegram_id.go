package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamUserRepository) GetUserTelegramID(id *uuid.UUID) (*int, *core.Status) {
	result := r.db.QueryRow("select telegram_id from users where id = ?", id.String())
	var telegramID int
	err := result.Scan(&telegramID)

	if err != nil {
		return nil, core.NotFoundError("Telegram user not found")
	}

	return &telegramID, core.StatusSuccess()
}
