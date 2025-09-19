package iam_permission_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *IamPermissionRepository) ReadPermission(name string) (*iam_domain.PermissionRead, *core.Status) {
	result := r.db.QueryRow("select name, description, meta, archived from permissions where name = ?", name)

	if result.Err() != nil {
		return nil, core.NotFoundError("iam.permission")
	}

	var permission iam_domain.PermissionRead
	err := result.Scan(&permission.Name, &permission.Description, &permission.Meta, &permission.Archived)
	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	return &permission, core.StatusSuccess()
}
