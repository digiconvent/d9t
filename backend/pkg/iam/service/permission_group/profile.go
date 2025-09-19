package iam_permission_group_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamPermissionGroupService) GetPermissionGroupProfile(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status) {
	profile := &iam_domain.PermissionGroupProfile{}
	group, status := s.repository.PermissionGroup.GetPermissionGroup(id)
	if status.Err() {
		return nil, status
	}

	profile.PermissionGroup = group
	users, status := s.repository.PermissionGroup.ListGroupUsers(id)
	if status.Err() {
		return nil, status
	}
	profile.Users = users

	permissionGroups, status := s.repository.PermissionGroup.ListPermissionGroupAncestors(id)
	if status.Err() {
		return nil, status
	}
	profile.Ancestors = permissionGroups

	permissionGroups, status = s.repository.PermissionGroup.ListPermissionGroupDescendants(id)
	if status.Err() {
		return nil, status
	}
	profile.Descendants = permissionGroups

	policies, status := s.repository.PermissionGroup.ListPermissionGroupPolicies(id)
	if status.Err() {
		return nil, status
	}
	profile.Policies = policies

	return profile, status
}
