package iam_policy_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *policyRepository) Create(policy *iam_domain.Policy) (*uuid.UUID, *core.Status) {
	id, _ := uuid.NewV7()

	query := `insert into policies (id, name, description, votes_required) values (?, ?, ?, ?)`

	_, err := r.db.Exec(query, id, policy.Name, policy.Description, policy.VotesRequired)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			if strings.Contains(err.Error(), "name") {
				return nil, core.ConflictError("policy name already exists")
			}
		}
		return nil, core.InternalError("failed to create policy")
	}

	return &id, core.StatusCreated()
}
