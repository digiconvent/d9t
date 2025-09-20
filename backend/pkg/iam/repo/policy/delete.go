package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *policyRepository) Delete(id *uuid.UUID) *core.Status {
	query := `delete from policies where id = ?`

	result, err := r.db.Exec(query, id.String())
	if err != nil {
		return core.InternalError("failed to delete policy")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check delete result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("policy not found")
	}

	return core.StatusSuccess()
}