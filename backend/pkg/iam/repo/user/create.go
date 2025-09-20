package iam_user_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) Create(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	id, _ := uuid.NewV7()

	query := `insert into users (id, email, first_name, last_name, password_hash, telegram, enabled, joined_at) values (?, ?, ?, ?, ?, ?, ?, datetime('now'))`

	_, err := r.db.Exec(query, id, user.Email, user.FirstName, user.LastName, user.PasswordHash, user.Telegram, user.Enabled)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			if strings.Contains(err.Error(), "email") {
				return nil, core.ConflictError("email already exists")
			}
		}
		return nil, core.InternalError("failed to create user")
	}

	return &id, core.StatusCreated()
}
