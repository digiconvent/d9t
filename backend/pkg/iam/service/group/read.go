package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *groupService) Read(id *uuid.UUID) (*iam_domain.Group, *core.Status) {
	return s.repo.Read(id)
}

func (s *groupService) ReadProfile(id *uuid.UUID) (*iam_domain.GroupProfile, *core.Status) {
	profile := &iam_domain.GroupProfile{}
	group, status := s.repo.Read(id)
	if status.Err() {
		return nil, status
	}
	profile.Group = group

	ascendants, status := s.repo.ReadAscendants(id)
	if status.Err() {
		return nil, status
	}
	descendants, status := s.repo.ReadDescendants(id)
	if status.Err() {
		return nil, status
	}
	profile.Ascendants = ascendants
	profile.Descendants = descendants

	return profile, core.StatusSuccess()
}

func (s *groupService) ReadProxies() ([]*iam_domain.GroupProxy, *core.Status) {
	return s.repo.ReadProxies()
}
