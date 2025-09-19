package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *IamAuthService) LoginUser(emailaddress string, rawPassword string) (*uuid.UUID, *core.Status) {
	userId, hashedPassword, status := s.repository.Auth.ReadCredentials(emailaddress)

	if status.Err() {
		return nil, status
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)) != nil {
		return nil, core.UnauthorizedError("Invalid credentials")
	}

	return userId, core.StatusSuccess()
}
