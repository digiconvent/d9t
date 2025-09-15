package iam_policy_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPolicyRepository) Update(id *uuid.UUID, arg *iam_domain.Policy) core.Status {
	result, err := r.db.Exec(`update policies set name = ?, description = ?, required_votes = ? where id = ?`, arg.Name, arg.Description, arg.Description, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("iam.policy")
	}

	return *core.StatusNoContent()
}
