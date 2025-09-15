package iam_permission_group_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) RemovePermissionFromPermissionGroup(permissionGroupId *uuid.UUID, permission string) core.Status {
	result, err := r.db.Exec("delete from permission_group_has_permission where permission_group = ? and permission = ?", permissionGroupId.String(), permission)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return *core.NotFoundError(err.Error())
	}

	if rowsAffected == 0 {
		return *core.NotFoundError("iam.permission_group.missing_permission")
	}

	return *core.StatusNoContent()
}
