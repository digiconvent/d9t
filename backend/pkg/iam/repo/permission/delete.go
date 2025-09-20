package iam_permission_repository

import (
	"github.com/digiconvent/d9t/core"
)

func (r *permissionRepository) Delete(permission string) *core.Status {
	query := `delete from permissions where permission = ?`

	result, err := r.db.Exec(query, permission)
	if err != nil {
		return core.InternalError("failed to delete permission")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check delete result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("permission not found")
	}

	return core.StatusSuccess()
}