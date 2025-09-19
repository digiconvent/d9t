package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	pagination "github.com/digiconvent/d9t/core/page"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamUserServiceInterface interface {
	Create(user *iam_domain.User) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.User, *core.Status)
	GetUserProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status)
	ListUsers(fs *iam_domain.UserFilterSort) (*pagination.Page[iam_domain.UserFacade], *core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.User) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status
	IsEnabled(id *uuid.UUID) (bool, *core.Status)
}

func (i *IamUserService) Create(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	id, status := i.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, status
	}
	return id, status
}

type IamUserService struct {
	repository *iam_repository.IamRepository
}

func NewUserService(iamRepo *iam_repository.IamRepository) *IamUserService {
	return &IamUserService{
		repository: iamRepo,
	}
}
