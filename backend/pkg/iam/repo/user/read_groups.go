package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) ReadGroups(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status) {
	query := `select g.id, g.name, g.type from groups g join group_has_user ghu on g.id = ghu."group" where ghu."user" = ?`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return nil, core.InternalError("failed to read user groups: " + err.Error())
	}
	defer rows.Close()

	var groups []*iam_domain.GroupProxy
	for rows.Next() {
		group := &iam_domain.GroupProxy{}
		err := rows.Scan(&group.Id, &group.Name, &group.Type)
		if err != nil {
			return nil, core.InternalError("failed to scan user group")
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate user groups")
	}

	return groups, core.StatusSuccess()
}