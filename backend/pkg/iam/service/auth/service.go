package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
)

type AuthServiceInterface interface {
	HashPassword(password string) (string, *core.Status)
	VerifyPassword(password, hash string) *core.Status
}

type authService struct{}

func NewAuthService() AuthServiceInterface {
	return &authService{}
}