package iam_group_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) Update(group *iam_domain.Group) *core.Status {
	query := `update groups set name = ?, type = ?, parent = ?, description = ? where id = ?`

	result, err := r.db.Exec(query, group.Name, group.Type, group.Parent, group.Description, group.Id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			if strings.Contains(err.Error(), "name") {
				return core.ConflictError("group name already exists")
			}
		}
		if strings.Contains(err.Error(), "would create cycle") {
			return core.BadRequestError("cannot create group hierarchy cycle")
		}
		return core.InternalError("failed to update group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check update result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("group not found")
	}

	return core.StatusSuccess()
}

func (r *groupRepository) SetParent(group, parent *uuid.UUID) *core.Status {
	query := `update groups set parent = ? where id = ?`

	var parentString *string
	if parent != nil {
		s := parent.String()
		parentString = &s
	}

	result, err := r.db.Exec(query, parentString, group.String())
	if err != nil {
		if strings.Contains(err.Error(), "would create cycle") {
			return core.BadRequestError("cannot create group hierarchy cycle")
		}
		return core.InternalError("failed to set group parent")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalError("failed to check update result")
	}

	if rowsAffected == 0 {
		return core.NotFoundError("group not found")
	}

	return core.StatusSuccess()
}
