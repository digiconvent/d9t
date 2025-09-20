package iam_user_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *userRepository) AddGroup(user, group *uuid.UUID) *core.Status {
	query := `insert into group_has_user ("group", "user", start_at) values (?, ?, datetime('now'))`

	_, err := r.db.Exec(query, group.String(), user.String())
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			return core.ConflictError("user is already a member of this group")
		}
		return core.InternalError("failed to add user to group")
	}

	return core.StatusCreated()
}
