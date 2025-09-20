package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *policyService) RemovePermission(policy *uuid.UUID, permission string) *core.Status {
	return s.repo.RemovePermission(policy, permission)
}