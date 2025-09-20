package iam_policy_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *policyRepository) Update(policy *iam_domain.Policy) *core.Status {
	query := `update policies set name = ?, description = ?, votes_required = ? where id = ?`

	result, err := r.db.Exec(query, policy.Name, policy.Description, policy.VotesRequired, policy.Id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			if strings.Contains(err.Error(), "name") {
				return core.ConflictError("policy name already exists")
			}
		}
		return core.InternalError("failed to update policy")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check update result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("policy not found")
	}

	return core.StatusSuccess()
}
