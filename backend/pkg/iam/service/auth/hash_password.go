package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) HashPassword(password string) (string, *core.Status) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", core.InternalError("failed to hash password")
	}

	return string(hash), core.StatusSuccess()
}