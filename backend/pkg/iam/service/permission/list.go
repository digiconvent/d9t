package iam_permission_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IamPermissionService) List() ([]*iam_domain.PermissionRead, *core.Status) {
	permissions, status := service.repository.Permission.ListPermissions()
	if status.Err() {
		return nil, &status
	}
	return permissions, &status
}
