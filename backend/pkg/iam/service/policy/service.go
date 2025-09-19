package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamPolicyServiceInterface interface {
	Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status)
	Update(id *uuid.UUID, policy *iam_domain.Policy) *core.Status
	Delete(id *uuid.UUID) *core.Status
}

type IamPolicyService struct {
	repository *iam_repository.IamRepository
}

func NewPolicyService(iamRepo *iam_repository.IamRepository) IamPolicyServiceInterface {
	return &IamPolicyService{
		repository: iamRepo,
	}
}
