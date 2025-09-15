package iam_auth_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
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

func (i *IamAuthService) Create(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status) {
	id, status := i.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}
	return id, &status
}

type IamAuthService struct {
	repository *iam_repository.IamRepository
}

func NewUserService(iamRepo *iam_repository.IamRepository) *IamAuthService {
	return &IamAuthService{
		repository: iamRepo,
	}
}
