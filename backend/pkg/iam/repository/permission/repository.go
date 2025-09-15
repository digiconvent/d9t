package iam_permission_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

type IamPermissionRepositoryInterface interface {
	CreatePermission(permission *iam_domain.PermissionWrite) core.Status
	DeletePermission(name string) core.Status
	GetPermission(name string) (*iam_domain.PermissionRead, core.Status)
	ListPermissionPermissionGroups(name string) ([]*iam_domain.PermissionGroupFacade, core.Status)
	ListPermissions() ([]*iam_domain.PermissionRead, core.Status)
	ListPermissionDescendants(name string) ([]*iam_domain.PermissionFacade, core.Status)
	ListPermissionUsers(name string) ([]*iam_domain.UserFacade, core.Status)
}

type IamPermissionRepository struct {
	db db.DatabaseInterface
}

func NewIamPermissionRepository(db db.DatabaseInterface) IamPermissionRepositoryInterface {
	return &IamPermissionRepository{
		db: db,
	}
}
