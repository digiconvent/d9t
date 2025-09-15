package iam_permission_group_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamPermissionGroupService) GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroup, *core.Status) {
	read, status := s.repository.PermissionGroup.GetPermissionGroup(id)

	if status.Err() {
		return nil, &status
	}

	// read.Policies, status = s.iamRepository.PermissionGroup.ListPermissionGroupPolicies(&read.Id)

	return read, &status
}
