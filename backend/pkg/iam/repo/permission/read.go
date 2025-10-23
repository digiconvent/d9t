package iam_permission_repository

import (
	"database/sql"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *permissionRepository) Read(permission string) (*iam_domain.Permission, *core.Status) {
	query := `select permission from permissions where permission = ?`

	row, err := r.db.QueryRow(query, permission)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NotFoundError("permission not found")
		}
		return nil, core.InternalError("failed to read permission")
	}

	perm := &iam_domain.Permission{}
	row.Scan(&perm.Permission)

	return perm, core.StatusSuccess()
}

func (r *permissionRepository) List() ([]*iam_domain.Permission, *core.Status) {
	query := `select permission from permissions`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, core.InternalError("failed to list permissions")
	}
	defer rows.Close()

	var permissions []*iam_domain.Permission
	for rows.Next() {
		perm := &iam_domain.Permission{}
		err := rows.Scan(&perm.Permission)
		if err != nil {
			return nil, core.InternalError("failed to scan permission")
		}
		permissions = append(permissions, perm)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate permissions")
	}

	return permissions, core.StatusSuccess()
}
