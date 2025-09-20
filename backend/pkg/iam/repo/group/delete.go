package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *groupRepository) Delete(id *uuid.UUID) *core.Status {
	query := `delete from groups where id = ?`

	result, err := r.db.Exec(query, id.String())
	if err != nil {
		return core.InternalError("failed to delete group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check delete result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("group not found")
	}

	return core.StatusSuccess()
}