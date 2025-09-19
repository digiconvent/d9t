package iam_permission_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (service *IamPermissionService) Profile(name string) (*iam_domain.PermissionProfile, *core.Status) {
	permission, status := service.repository.Permission.ReadPermission(name)
	if status.Err() {
		return nil, status
	}

	descendantPermissions, status := service.repository.Permission.ListPermissionDescendants(name)
	if status.Err() {
		return nil, status
	}

	policies, status := service.repository.Permission.ListPolicies(name)

	return &iam_domain.PermissionProfile{
		Permission:  permission,
		Policies:    policies,
		Descendants: descendantPermissions,
	}, core.StatusSuccess()
}
