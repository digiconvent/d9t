package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IamPolicyRepository) Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status) {
	result := r.db.QueryRow("select name, description, required_votes from policies where id = ?", id)

	if result.Err() != nil {
		return nil, core.NotFoundError("iam.policy")
	}

	var policy = iam_domain.Policy{
		Id: id,
	}
	err := result.Scan(&policy.Name, &policy.Description, &policy.RequiredVotes)
	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	return &policy, core.StatusSuccess()
}
