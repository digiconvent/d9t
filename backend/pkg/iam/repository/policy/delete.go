package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *IamPolicyRepository) Delete(id *uuid.UUID) *core.Status {
	result, err := r.db.Exec("delete from policies where id = ?", id)

	if err != nil {
		return core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return core.NotFoundError("Policy not found")
	}

	return core.StatusNoContent()
}
