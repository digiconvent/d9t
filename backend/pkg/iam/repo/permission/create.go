package iam_permission_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *permissionRepository) Create(permission *iam_domain.Permission) *core.Status {
	query := `insert into permissions (permission) values (?)`

	_, err := r.db.Exec(query, permission.Permission)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			return core.ConflictError("permission already exists")
		}
		return core.InternalError("failed to create permission")
	}

	return core.StatusCreated()
}
