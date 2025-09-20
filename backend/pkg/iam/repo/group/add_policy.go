package iam_group_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *groupRepository) AddPolicy(group, policy *uuid.UUID) *core.Status {
	query := `insert into group_has_policy ("group", policy) values (?, ?)`

	_, err := r.db.Exec(query, group, policy)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			return core.ConflictError("group already has this policy")
		}
		return core.InternalError("failed to add policy to group: " + err.Error())
	}

	return core.StatusCreated()
}
