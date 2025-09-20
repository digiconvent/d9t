package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/pkg/iam/repo/group"
	"github.com/google/uuid"
)

type GroupServiceInterface interface {
	Create(group *iam_domain.Group) (*uuid.UUID, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Group, *core.Status)
	ReadProfile(id *uuid.UUID) (*iam_domain.GroupProfile, *core.Status)
	ReadProxies() ([]*iam_domain.GroupProxy, *core.Status)
	Update(group *iam_domain.Group) *core.Status
	Delete(id *uuid.UUID) *core.Status
	AddUser(group, user *uuid.UUID) *core.Status
	RemoveUser(group, user *uuid.UUID) *core.Status
	AddPolicy(group, policy *uuid.UUID) *core.Status
	RemovePolicy(group, policy *uuid.UUID) *core.Status
	SetParent(group, parent *uuid.UUID) *core.Status
}

type groupService struct {
	repo iam_group_repository.GroupRepositoryInterface
}

func NewGroupService(repo iam_group_repository.GroupRepositoryInterface) GroupServiceInterface {
	return &groupService{repo: repo}
}