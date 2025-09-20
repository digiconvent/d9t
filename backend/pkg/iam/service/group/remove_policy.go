package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *groupService) RemovePolicy(group, policy *uuid.UUID) *core.Status {
	return s.repo.RemovePolicy(group, policy)
}