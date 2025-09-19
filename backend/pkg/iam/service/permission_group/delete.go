package iam_permission_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *IamPermissionGroupService) DeletePermissionGroup(id *uuid.UUID) *core.Status {
	status := s.repository.PermissionGroup.DeletePermissionGroup(id)
	return status
}
