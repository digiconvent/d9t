package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *policyRepository) ReadPermissions(id *uuid.UUID) ([]string, *core.Status) {
	query := `select permission from policy_has_permission where policy = ?`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return nil, core.InternalError("failed to read policy permissions: " + err.Error())
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var permission string
		err := rows.Scan(&permission)
		if err != nil {
			return nil, core.InternalError("failed to scan policy permission")
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate policy permissions")
	}

	return permissions, core.StatusSuccess()
}
