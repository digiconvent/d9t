package iam_policy_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *policyRepository) AddPermission(policy *uuid.UUID, permission string) *core.Status {
	query := `insert into policy_has_permission (policy, permission) values (?, ?)`

	_, err := r.db.Exec(query, policy.String(), permission)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			return core.ConflictError("policy already has this permission")
		}
		return core.InternalError("failed to add permission to policy")
	}

	return core.StatusCreated()
}
