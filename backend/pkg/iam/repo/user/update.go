package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) Update(user *iam_domain.User) *core.Status {
	query := `update users set email = ?, first_name = ?, last_name = ?, password_hash = ?, telegram = ?, enabled = ? where id = ?`

	result, err := r.db.Exec(query, user.Email, user.FirstName, user.LastName, user.PasswordHash, user.Telegram, user.Enabled, user.Id)
	if err != nil {
		return core.InternalError("failed to update user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check update result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("user not found")
	}

	return core.StatusSuccess()
}

func (r *userRepository) SetEnabled(id *uuid.UUID, enabled bool) *core.Status {
	query := `update users set enabled = ? where id = ?`

	result, err := r.db.Exec(query, enabled, id.String())
	if err != nil {
		return core.InternalError("failed to update user enabled status")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check update result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("user not found")
	}

	return core.StatusSuccess()
}
