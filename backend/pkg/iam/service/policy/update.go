package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
)

func (s *policyService) Update(policy *iam_domain.Policy) *core.Status {
	return s.repo.Update(policy)
}