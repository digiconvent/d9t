package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/pkg/iam/repo/policy"
	"github.com/google/uuid"
)

type PolicyServiceInterface interface {
	Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status)
	ReadProfile(id *uuid.UUID) (*iam_domain.PolicyProfile, *core.Status)
	ReadProxies() ([]*iam_domain.PolicyProxy, *core.Status)
	Update(policy *iam_domain.Policy) *core.Status
	Delete(id *uuid.UUID) *core.Status
	AddPermission(policy *uuid.UUID, permission string) *core.Status
	RemovePermission(policy *uuid.UUID, permission string) *core.Status
}

type policyService struct {
	repo iam_policy_repository.PolicyRepositoryInterface
}

func NewPolicyService(repo iam_policy_repository.PolicyRepositoryInterface) PolicyServiceInterface {
	return &policyService{repo: repo}
}