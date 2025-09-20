package iam_auth_repository

import (
	"database/sql"

	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *authRepository) GetPasswordHash(user *uuid.UUID) (string, *core.Status) {
	query := `select password_hash from users where id = ?`

	var hash string
	err := r.db.QueryRow(query, user.String()).Scan(&hash)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", core.NotFoundError("user not found")
		}
		return "", core.InternalError("failed to get password hash")
	}

	return hash, core.StatusSuccess()
}
