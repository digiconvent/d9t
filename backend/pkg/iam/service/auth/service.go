package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamAuthServiceInterface interface {
	LoginTelegramUser(body, botToken string) (*uuid.UUID, *core.Status)
	ConnectTelegramUser(body, botToken string, userId *uuid.UUID) *core.Status

	ResetPassword(emailaddress string) (string, *core.Status)
	SetUserPassword(id *uuid.UUID, password string) *core.Status
	LoginUser(emailaddress, password string) (*uuid.UUID, *core.Status)

	GenerateJwt(userId *uuid.UUID) (string, *core.Status)
	VerifyJwt(token string) (*uuid.UUID, *core.Status)
}

func (i *IamAuthService) Create(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	id, status := i.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, status
	}
	return id, status
}

type IamAuthService struct {
	repository *iam_repository.IamRepository
}

func (s *IamAuthService) ConnectTelegramUser(body string, botToken string, userId *uuid.UUID) *core.Status {
	telegramId, status := s.repository.Auth.GetTelegramId(body, botToken)
	if status.Err() {
		return status
	}

	status = s.repository.User.RegisterTelegramUser(*telegramId, userId)
	if status.Err() {
		return status
	}

	return status
}

func NewAuthService(repo *iam_repository.IamRepository) IamAuthServiceInterface {
	return &IamAuthService{
		repository: repo,
	}
}
