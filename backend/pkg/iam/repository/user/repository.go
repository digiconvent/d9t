package iam_user_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/pagination"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

type IamUserRepositoryInterface interface {
	CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, core.Status)
	GetUserByEmailaddress(emailaddress string) (*iam_domain.UserRead, core.Status)
	GetUserByID(id *uuid.UUID) (*iam_domain.UserRead, core.Status)
	ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status)
	List(*iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], core.Status)
	ListUserGroups(userId *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	RegisterTelegramUser(telegramId int, userId *uuid.UUID) core.Status
	SetEnabled(id *uuid.UUID, enabled bool) core.Status
	IsEnabled(id *uuid.UUID) (bool, core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status
	UserHasPermission(userId *uuid.UUID, permission string) bool
	GetUserByTelegramID(id *int) (*uuid.UUID, core.Status)
	GetTelegramID(dataString, botToken string) (*int, core.Status)
	GetUserTelegramID(id *uuid.UUID) (*int, core.Status)
}

type IamUserRepository struct {
	db db.DatabaseInterface
}

func NewIamUserRepository(db db.DatabaseInterface) IamUserRepositoryInterface {
	return &IamUserRepository{
		db: db,
	}
}
