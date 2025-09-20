package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) ReadAscendants(id *uuid.UUID) ([]*iam_domain.GroupProxy, *core.Status) {
	query := `select id, name, type from groups where id in (select ancestor from group_has_ancestors where "id" = ?)`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, core.InternalError("failed to read group ascendants")
	}
	defer rows.Close()

	var ascendants []*iam_domain.GroupProxy
	for rows.Next() {
		ascendant := &iam_domain.GroupProxy{}
		err := rows.Scan(&ascendant.Id, &ascendant.Name, &ascendant.Type)
		if err != nil {
			return nil, core.InternalError("failed to scan group ascendant")
		}
		ascendants = append(ascendants, ascendant)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate group ascendants")
	}

	return ascendants, core.StatusSuccess()
}
