package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, *core.Status) {
	var permissionGroups []*iam_domain.PermissionGroupFacade
	rows, err := r.db.Query(`select id, name, abbr, meta, parent from permission_groups`)

	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		permissionGroup := iam_domain.PermissionGroupFacade{}

		var parentId *string

		err = rows.Scan(
			&permissionGroup.Id,
			&permissionGroup.Name,
			&permissionGroup.Abbr,
			&permissionGroup.Meta,
			&parentId,
		)

		if parentId != nil {
			parsedParentId, err := uuid.Parse(*parentId)
			if err == nil {
				permissionGroup.Parent = &parsedParentId
			}
		}

		if err != nil {
			return nil, core.InternalError(err.Error())
		}
		permissionGroups = append(permissionGroups, &permissionGroup)
	}
	return permissionGroups, core.StatusSuccess()
}
