package iam_permission_group_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamPermissionGroupService) UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroup) *core.Status {
	if arg == nil {
		return core.UnprocessableContentError("iam.permission_group.update.missing_data")
	}
	if arg.Name == "" {
		return core.UnprocessableContentError("iam.permission_group.update.invalid_name")
	}

	// status := s.iamRepository.PermissionGroup.SetPermissionsForPermissionGroup(id, arg.Policies)

	// if status.Err() {
	// 	return &status
	// }

	status := s.repository.PermissionGroup.UpdatePermissionGroup(id, arg)

	return &status
}
