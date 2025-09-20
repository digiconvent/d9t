package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *policyRepository) ReadGroups(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status) {
	query := `select g.id, g.name from group_has_policy ghp left join groups g on ghp."group" = g.id where ghp.policy = ?`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, core.InternalError("failed to read groups")
	}
	defer rows.Close()

	var groups []*iam_domain.GroupProxy
	for rows.Next() {
		group := &iam_domain.GroupProxy{}
		err := rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return nil, core.InternalError("failed to scan group")
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate groups")
	}

	return groups, core.StatusSuccess()
}
