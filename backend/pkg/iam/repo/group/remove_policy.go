package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *groupRepository) RemovePolicy(group, policy *uuid.UUID) *core.Status {
	query := `delete from group_has_policy where "group" = ? and policy = ?`

	result, err := r.db.Exec(query, group.String(), policy.String())
	if err != nil {
		return core.InternalError("failed to remove policy from group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check removal result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("group does not have this policy")
	}

	return core.StatusSuccess()
}
