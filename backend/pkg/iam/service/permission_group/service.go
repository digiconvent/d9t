package iam_permission_group_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IamPermissionGroupServiceInterface interface {
	CreatePermissionGroup(arg *iam_domain.PermissionGroup) (*uuid.UUID, *core.Status)

	GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroup, *core.Status)
	GetPermissionGroupProfile(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, *core.Status)

	AddUserToPermissionGroup(permissionGroup *uuid.UUID, userId *uuid.UUID) *core.Status
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroup) *core.Status

	DeletePermissionGroup(id *uuid.UUID) *core.Status
}

type IamPermissionGroupService struct {
	repository *iam_repository.IamRepository
}

func NewIamPermissionGroupService(iamRepository *iam_repository.IamRepository) IamPermissionGroupServiceInterface {
	return &IamPermissionGroupService{
		repository: iamRepository,
	}
}
