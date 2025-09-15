package iam_permission_group_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) CreatePermissionGroup(arg *iam_domain.PermissionGroup) (*uuid.UUID, core.Status) {
	id, _ := uuid.NewV7()

	_, err := r.db.Exec(`insert into permission_groups (id, name, abbr, description, parent) values (?, ?, ?, ?, ?)`, id, arg.Name, arg.Abbr, arg.Description, arg.Parent)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return &id, *core.StatusCreated()
}
