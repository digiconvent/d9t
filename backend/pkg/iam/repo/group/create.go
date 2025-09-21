package iam_group_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) Create(group *iam_domain.Group) (*uuid.UUID, *core.Status) {
	id, _ := uuid.NewV7()

	if group.Parent == nil {
		return nil, core.UnprocessableContentError("parent cannot be empty")
	}

	query := `insert into groups (id, name, type, parent, description) values (?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, id, group.Name, group.Type, group.Parent, group.Description)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") {
			if strings.Contains(err.Error(), "name") {
				return nil, core.ConflictError("group name already exists")
			}
		}
		if strings.Contains(err.Error(), "would create cycle") {
			return nil, core.ConflictError("cannot create group hierarchy cycle")
		}
		return nil, core.InternalError("failed to create group: " + err.Error())
	}

	return &id, core.StatusCreated()
}
