package iam_user_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) Create(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	id, _ := uuid.NewV7()

	query := `insert into users (id, email, first_name, last_name, telegram, enabled, joined_at) values (?, ?, ?, ?, ?, ?, datetime('now'))`

	_, err := r.db.ExecDebug(query, id, user.Email, user.FirstName, user.LastName, user.Telegram, false)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			if strings.Contains(err.Error(), "email") {
				return nil, core.ConflictError("iam.user.email.duplicate")
			}
		}
		return nil, core.InternalError("failed to create user")
	}

	return &id, core.StatusCreated()
}
