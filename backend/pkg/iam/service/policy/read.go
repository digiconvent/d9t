package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *policyService) Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status) {
	return s.repo.Read(id)
}

func (s *policyService) ReadProfile(id *uuid.UUID) (*iam_domain.PolicyProfile, *core.Status) {
	policy, status := s.repo.Read(id)
	if status.Err() {
		return nil, status
	}
	profile := &iam_domain.PolicyProfile{
		Policy: policy,
	}
	groups, status := s.repo.ReadGroups(id)
	if status.Err() {
		return nil, status
	}
	profile.Groups = groups

	permissions, status := s.repo.ReadPermissions(id)
	if status.Err() {
		return nil, status
	}
	profile.Permissions = permissions

	return profile, core.StatusSuccess()
}

func (s *policyService) ReadProxies() ([]*iam_domain.PolicyProxy, *core.Status) {
	return s.repo.ReadProxies()
}
