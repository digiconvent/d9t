package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) ReadDescendants(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status) {
	query := `select id, name, type from groups where parent = ?`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return nil, core.InternalError("failed to read group descendants")
	}
	defer rows.Close()

	var descendants []*iam_domain.GroupProxy
	for rows.Next() {
		descendant := &iam_domain.GroupProxy{}
		err := rows.Scan(&descendant.Id, &descendant.Name, &descendant.Type)
		if err != nil {
			return nil, core.InternalError("failed to scan group descendant")
		}
		descendants = append(descendants, descendant)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate group descendants")
	}

	return descendants, core.StatusSuccess()
}
