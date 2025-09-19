package iam_permission_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *IamPermissionGroupService) AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) *core.Status {
	status := s.repository.PermissionGroup.AddUserToPermissionGroup(permissionGroup, userId)
	return status
}
