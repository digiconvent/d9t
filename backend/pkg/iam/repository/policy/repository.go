package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

type IamPolicyRepositoryInterface interface {
	Create(data *iam_domain.Policy) (*uuid.UUID, *core.Status)
	List() ([]*iam_domain.Policy, *core.Status)
	Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status)
	Update(id *uuid.UUID, data *iam_domain.Policy) *core.Status
	Delete(id *uuid.UUID) *core.Status
}

type IamPolicyRepository struct {
	db db.DatabaseInterface
}

func NewIamPolicyRepository(db db.DatabaseInterface) IamPolicyRepositoryInterface {
	return &IamPolicyRepository{
		db: db,
	}
}
