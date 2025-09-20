package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *groupService) AddPolicy(group, policy *uuid.UUID) *core.Status {
	return s.repo.AddPolicy(group, policy)
}