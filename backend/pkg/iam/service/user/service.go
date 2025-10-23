package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_user_repository "github.com/digiconvent/d9t/pkg/iam/repo/user"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	Create(user *iam_domain.User) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.User, *core.Status)
	ReadProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status)
	ReadProxies() ([]*iam_domain.UserProxy, *core.Status)
	ReadByEmail(email string) (*iam_domain.User, *core.Status)
	Update(user *iam_domain.User) *core.Status
	Delete(id *uuid.UUID) *core.Status
	AddGroup(user, group *uuid.UUID) *core.Status
	RemoveGroup(user, group *uuid.UUID) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status
	UserHasPermission(user *uuid.UUID, permission string) ([]*iam_domain.PolicyProxy, *core.Status)
}

type userService struct {
	repo iam_user_repository.UserRepositoryInterface
}

func NewUserService(repo iam_user_repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}
