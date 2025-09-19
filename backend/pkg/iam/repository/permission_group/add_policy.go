package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) AddPolicyToPermissionGroup(policyId *uuid.UUID, permissionGroupId *uuid.UUID) *core.Status {
	if policyId == nil || permissionGroupId == nil {
		return core.UnprocessableContentError("permission group and user id must be provided")
	}

	res, err := r.db.Exec(`insert into policy_has_permission_group (policy, permission_group) values (?,?)`, policyId, permissionGroupId)
	if err != nil {
		return core.InternalError(err.Error())
	}
	n, err := res.RowsAffected()
	if err != nil {
		return core.InternalError(err.Error())
	}
	if n == 0 {
		return core.ConflictError("")
	}
	return core.StatusSuccess()
}
