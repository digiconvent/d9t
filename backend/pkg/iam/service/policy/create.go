package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (i *IamPolicyService) Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status) {
	return i.repository.Policy.Create(policy)
}
