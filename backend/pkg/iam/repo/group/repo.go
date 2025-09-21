package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

type GroupRepositoryInterface interface {
	Create(group *iam_domain.Group) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Group, *core.Status)
	ReadProxies() ([]*iam_domain.GroupProxy, *core.Status)
	ReadAscendants(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status)
	ReadDescendants(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status)
	ReadUsers(id *uuid.UUID) ([]*iam_domain.UserProxy, *core.Status)
	ReadPolicies(id *uuid.UUID) ([]*iam_domain.PolicyProxy, *core.Status)
	Update(group *iam_domain.Group) *core.Status
	Delete(id *uuid.UUID) *core.Status
	AddUser(group, user *uuid.UUID) *core.Status
	RemoveUser(group, user *uuid.UUID) *core.Status
	AddPolicy(group, policy *uuid.UUID) *core.Status
	RemovePolicy(group, policy *uuid.UUID) *core.Status
	SetParent(group, parent *uuid.UUID) *core.Status
}

type groupRepository struct {
	db db.DatabaseInterface
}

func NewGroupRepository(database db.DatabaseInterface) GroupRepositoryInterface {
	return &groupRepository{db: database}
}
