package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (i *IamPolicyService) Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status) {
	return i.repository.Policy.Read(id)
}
