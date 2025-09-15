package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamUserServiceInterface interface {
	Create(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.UserRead, *core.Status)
	GetUserProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status)
	ListUsers(fs *iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], *core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status
	IsEnabled(id *uuid.UUID) (bool, *core.Status)
}

func (i *IamUserService) Create(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status) {
	id, status := i.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}
	return id, &status
}

type IamUserService struct {
	repository *iam_repository.IamRepository
}

func NewUserService(iamRepo *iam_repository.IamRepository) *IamUserService {
	return &IamUserService{
		repository: iamRepo,
	}
}
