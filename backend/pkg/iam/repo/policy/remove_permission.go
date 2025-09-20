package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *policyRepository) RemovePermission(policy *uuid.UUID, permission string) *core.Status {
	query := `delete from policy_has_permission where policy = ? and permission = ?`

	result, err := r.db.Exec(query, policy.String(), permission)
	if err != nil {
		return core.InternalError("failed to remove permission from policy")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check removal result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("policy does not have this permission")
	}

	return core.StatusSuccess()
}
