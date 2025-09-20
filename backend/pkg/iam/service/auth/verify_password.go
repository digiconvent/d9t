package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) VerifyPassword(password, hash string) *core.Status {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return core.UnauthorizedError("invalid password")
	}

	return core.StatusSuccess()
}