package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *policyService) Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status) {
	return s.repo.Create(policy)
}