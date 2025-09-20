package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *groupRepository) RemoveUser(group, user *uuid.UUID) *core.Status {
	query := `delete from group_has_user where "group" = ? and "user" = ? and start_at = (
			select start_at from current_group_memberships
			where "group" = ? and "user" = ?
		)`

	result, err := r.db.Exec(query, group.String(), user.String(), group.String(), user.String())
	if err != nil {
		return core.InternalError("failed to remove user from group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check removal result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("user is not a member of this group")
	}

	return core.StatusSuccess()
}
