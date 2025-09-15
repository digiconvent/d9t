package iam_permission_repository

import (
	"github.com/DigiConvent/testd9t/core"
)

func (r *IamPermissionRepository) DeletePermission(name string) core.Status {
	result, err := r.db.Exec("delete from permissions where name = ?", name)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("Permission not found")
	}

	return *core.StatusNoContent()
}
