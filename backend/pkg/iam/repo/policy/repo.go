package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

type PolicyRepositoryInterface interface {
	Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status)

	Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status)
	ReadGroups(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status)
	ReadPermissions(id *uuid.UUID) ([]string, *core.Status)
	ReadProxies() ([]*iam_domain.PolicyProxy, *core.Status)

	Update(policy *iam_domain.Policy) *core.Status

	Delete(id *uuid.UUID) *core.Status

	AddPermission(policy *uuid.UUID, permission string) *core.Status
	RemovePermission(policy *uuid.UUID, permission string) *core.Status
}

type policyRepository struct {
	db db.DatabaseInterface
}

func NewPolicyRepository(database db.DatabaseInterface) PolicyRepositoryInterface {
	return &policyRepository{db: database}
}
