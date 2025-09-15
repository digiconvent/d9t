package iam_permission_group_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

type IamPermissionGroupRepository struct {
	db db.DatabaseInterface
}

type IamPermissionGroupRepositoryInterface interface {
	AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) core.Status
	CreatePermissionGroup(arg *iam_domain.PermissionGroup) (*uuid.UUID, core.Status)
	DeletePermissionGroup(arg *uuid.UUID) core.Status
	GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroup, core.Status)
	ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	ListPermissionGroupAncestors(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	ListPermissionGroupDescendants(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	ListPermissionGroupPolicies(arg *uuid.UUID) ([]*iam_domain.Policy, core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, core.Status)
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) core.Status
	SetPermissionsForPermissionGroup(permissionGroupId *uuid.UUID, permissions []string) core.Status
	AddPermissionToPermissionGroup(permissionGroupId *uuid.UUID, permission string) core.Status
	RemovePermissionFromPermissionGroup(permissionGroupId *uuid.UUID, permission string) core.Status
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroup) core.Status
}

func NewIamPermissionGroupRepository(db db.DatabaseInterface) IamPermissionGroupRepositoryInterface {
	return &IamPermissionGroupRepository{
		db: db,
	}
}
