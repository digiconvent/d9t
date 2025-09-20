package iam_permission_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
)

type PermissionRepositoryInterface interface {
	Create(permission *iam_domain.Permission) *core.Status
	Read(permission string) (*iam_domain.Permission, *core.Status)
	List() ([]*iam_domain.Permission, *core.Status)
	Delete(permission string) *core.Status
}

type permissionRepository struct {
	db db.DatabaseInterface
}

func NewPermissionRepository(database db.DatabaseInterface) PermissionRepositoryInterface {
	return &permissionRepository{db: database}
}
