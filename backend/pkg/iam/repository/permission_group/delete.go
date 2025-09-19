package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) DeletePermissionGroup(arg *uuid.UUID) *core.Status {
	if arg == nil {
		return core.UnprocessableContentError("ID is required")
	}
	res, err := r.db.Exec(`delete from permission_groups where id = ?`, arg.String())
	if err != nil {
		return core.InternalError(err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return core.NotFoundError("permission group not found")
	}
	return core.StatusNoContent()
}
