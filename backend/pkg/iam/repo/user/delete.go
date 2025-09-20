package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *userRepository) Delete(id *uuid.UUID) *core.Status {
	query := `delete from users where id = ?`

	result, err := r.db.Exec(query, id.String())
	if err != nil {
		return core.InternalError("failed to delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check delete result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("user not found")
	}

	return core.StatusSuccess()
}