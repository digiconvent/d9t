package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	uuid "github.com/google/uuid"
)

type IamPermissionGroupRepository struct {
	db db.DatabaseInterface
}

type IamPermissionGroupRepositoryInterface interface {
	AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) *core.Status
	AddPolicyToPermissionGroup(policyId, permissionGroupId *uuid.UUID) *core.Status
	CreatePermissionGroup(arg *iam_domain.PermissionGroup) (*uuid.UUID, *core.Status)
	DeletePermissionGroup(arg *uuid.UUID) *core.Status
	GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroup, *core.Status)
	ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, *core.Status)
	ListPermissionGroupAncestors(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, *core.Status)
	ListPermissionGroupDescendants(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, *core.Status)
	ListPermissionGroupPolicies(arg *uuid.UUID) ([]*iam_domain.PolicyFacade, *core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, *core.Status)
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroup) *core.Status
}

func NewIamPermissionGroupRepository(db db.DatabaseInterface) IamPermissionGroupRepositoryInterface {
	return &IamPermissionGroupRepository{
		db: db,
	}
}
