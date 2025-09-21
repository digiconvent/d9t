package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Create(user *iam_domain.User) (*uuid.UUID, *core.Status)

	Read(id *uuid.UUID) (*iam_domain.User, *core.Status)
	ReadProxies() ([]*iam_domain.UserProxy, *core.Status)
	ReadByEmail(email string) (*iam_domain.User, *core.Status)
	ReadGroups(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status)
	ReadPoliciesWithPermission(user *uuid.UUID, permission string) ([]*iam_domain.PolicyProxy, *core.Status)

	Update(user *iam_domain.User) *core.Status

	Delete(id *uuid.UUID) *core.Status

	AddGroup(user, group *uuid.UUID) *core.Status
	RemoveGroup(user, group *uuid.UUID) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status
}

type userRepository struct {
	db db.DatabaseInterface
}

func NewUserRepository(database db.DatabaseInterface) UserRepositoryInterface {
	return &userRepository{db: database}
}
