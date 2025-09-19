package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) *core.Status {
	if permissionGroup == nil || userId == nil {
		return core.UnprocessableContentError("permission group and user id must be provided")
	}

	_, err := r.db.Exec(`insert into permission_group_has_user (permission_group, user) values (?, ?)`, permissionGroup.String(), userId.String())

	if err != nil {
		return core.InternalError(err.Error())
	}

	return core.StatusSuccess()
}
