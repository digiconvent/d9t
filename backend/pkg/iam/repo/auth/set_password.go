package iam_auth_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"golang.org/x/crypto/bcrypt"
)

func (r *authRepository) SetPassword(userPassword *iam_domain.UserPassword) *core.Status {
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword.Password), bcrypt.DefaultCost)
	if err != nil {
		return core.InternalError("failed to hash password")
	}

	query := `update users set password_hash = ? where id = ?`

	result, err := r.db.Exec(query, string(hash), userPassword.User.String())
	if err != nil {
		return core.InternalError("failed to set password")
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
