package iam_permission_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
)

type IamPermissionServiceInterface interface {
	Create(data *iam_domain.PermissionWrite) *core.Status
	// Read(name string) (*iam_domain.PermissionRead, *core.Status)
	List() ([]*iam_domain.PermissionRead, *core.Status)
	Profile(name string) (*iam_domain.PermissionProfile, *core.Status)
	// Update(name string, data *iam_domain.PermissionWrite) *core.Status
	Delete(name string) *core.Status
}

type IamPermissionService struct {
	repository *iam_repository.IamRepository
}

func NewIamPermissionService(iamRepo *iam_repository.IamRepository) IamPermissionServiceInterface {
	return &IamPermissionService{
		repository: iamRepo,
	}
}
