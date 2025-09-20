package iam_permission_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_permission_repository "github.com/digiconvent/d9t/pkg/iam/repo/permission"
)

type PermissionServiceInterface interface {
	Create(permission *iam_domain.Permission) *core.Status
	Read(permission string) (*iam_domain.Permission, *core.Status)
	List() ([]*iam_domain.Permission, *core.Status)
	Delete(permission string) *core.Status
}

type permissionService struct {
	repo iam_permission_repository.PermissionRepositoryInterface
}

func NewPermissionService(repo iam_permission_repository.PermissionRepositoryInterface) PermissionServiceInterface {
	return &permissionService{repo: repo}
}
