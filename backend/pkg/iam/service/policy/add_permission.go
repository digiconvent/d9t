package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *policyService) AddPermission(policy *uuid.UUID, permission string) *core.Status {
	return s.repo.AddPermission(policy, permission)
}