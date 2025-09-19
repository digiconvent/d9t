package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IamPolicyRepository) Create(data *iam_domain.Policy) (*uuid.UUID, *core.Status) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, &core.Status{Code: 500, Message: err.Error()}
	}
	_, err = r.db.Exec(`insert into policies (id, name, description, required_votes) values (?,?,?,?)`, id, data.Name, data.Description, data.RequiredVotes)
	if err != nil {
		return nil, &core.Status{Code: 500, Message: err.Error()}
	}

	return &id, &core.Status{Code: 201}
}
