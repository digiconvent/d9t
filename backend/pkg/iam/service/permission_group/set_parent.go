package iam_permission_group_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IamPermissionGroupService) SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status {
	pg, status := s.repository.PermissionGroup.GetPermissionGroup(arg.Parent)
	if status.Err() {
		return &status
	}
	if pg.Meta == "role" {
		return core.UnprocessableContentError("cannot add a permission group to a role")
	}

	if pg.Meta == "status" {
		return core.UnprocessableContentError("cannot add a permission group to a status")
	}

	status = s.repository.PermissionGroup.SetParentPermissionGroup(arg)
	return &status
}
