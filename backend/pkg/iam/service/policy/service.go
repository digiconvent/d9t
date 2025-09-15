package iam_policy_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamPolicyServiceInterface interface {
	Create(user *iam_domain.Policy) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status)
	Update(id *uuid.UUID, user *iam_domain.Policy) *core.Status
	Delete(id *uuid.UUID) *core.Status
}

func (i *IamPolicyService) Create(user *iam_domain.Policy) (*uuid.UUID, *core.Status) {
	id, status := i.repository.Policy.Create(user)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}
	return id, &status
}

type IamPolicyService struct {
	repository *iam_repository.IamRepository
}

func NewUserService(iamRepo *iam_repository.IamRepository) *IamPolicyService {
	return &IamPolicyService{
		repository: iamRepo,
	}
}
