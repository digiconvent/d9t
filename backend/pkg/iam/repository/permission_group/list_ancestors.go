package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) ListPermissionGroupAncestors(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, *core.Status) {
	var permissionGroups = make([]*iam_domain.PermissionGroupFacade, 0)
	rows, err := r.db.Query(`select id, name, parent, implied from permission_group_has_permission_group_ancestors where root = ?`, arg.String())

	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var permissionGroup iam_domain.PermissionGroupFacade
		err := rows.Scan(&permissionGroup.Id, &permissionGroup.Name, &permissionGroup.Parent, &permissionGroup.Implied)
		if err != nil {
			return nil, core.InternalError(err.Error())
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, core.StatusSuccess()
}
