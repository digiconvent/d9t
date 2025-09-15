package iam_policy_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IamPolicyRepository) Create(data *iam_domain.Policy) (*uuid.UUID, core.Status) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, core.Status{Code: 500, Message: err.Error()}
	}
	_, err = r.db.Exec(`insert into policies (id, name, description, votes_required) values (?,?,?,?,?)`, id, data.Name, data.Description, data.RequiredVotes)
	if err != nil {
		return nil, core.Status{Code: 500, Message: err.Error()}
	}

	return &id, core.Status{Code: 201}
}
