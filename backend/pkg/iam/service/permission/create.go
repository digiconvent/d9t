package iam_permission_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IamPermissionService) Create(permission *iam_domain.PermissionWrite) *core.Status {
	status := service.repository.Permission.CreatePermission(permission)
	return &status
}
