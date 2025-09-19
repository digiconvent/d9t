package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamUserRepository) RegisterTelegramUser(telegramId int, userId *uuid.UUID) *core.Status {
	result, err := r.db.Exec("update users set telegram_id = ? where id = ?", telegramId, userId.String())
	if err != nil {
		return core.InternalError("Failed to update user: " + err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return core.InternalError("Failed to update user")
	}

	return core.StatusSuccess()

}
