package iam_permission_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
)

type IamPermissionRepositoryInterface interface {
	CreatePermission(permission *iam_domain.PermissionWrite) *core.Status

	ReadPermission(name string) (*iam_domain.PermissionRead, *core.Status)
	ListPermissions() ([]*iam_domain.PermissionRead, *core.Status)
	ListPolicies(name string) ([]*iam_domain.PolicyFacade, *core.Status)
	ListPermissionDescendants(name string) ([]*iam_domain.PermissionFacade, *core.Status)

	DeletePermission(name string) *core.Status
}

type IamPermissionRepository struct {
	db db.DatabaseInterface
}

func NewIamPermissionRepository(db db.DatabaseInterface) IamPermissionRepositoryInterface {
	return &IamPermissionRepository{
		db: db,
	}
}
