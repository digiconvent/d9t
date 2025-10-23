package iam_group_repository

import (
	"database/sql"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) Read(id *uuid.UUID) (*iam_domain.Group, *core.Status) {
	query := `select id, name, type, parent, description from groups where id = ?`

	row, err := r.db.QueryRow(query, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NotFoundError("group not found")
		}
		return nil, core.InternalError("failed to read group")
	}

	group := &iam_domain.Group{}
	err = row.Scan(
		&group.Id,
		&group.Name,
		&group.Type,
		&group.Parent,
		&group.Description,
	)
	if err != nil {
		return nil, core.NotFoundError("iam.permission_group.not_found")
	}

	return group, core.StatusSuccess()
}

func (r *groupRepository) ReadProxies() ([]*iam_domain.GroupProxy, *core.Status) {
	query := `select id, name, type from groups`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, core.InternalError("failed to read groups")
	}
	defer rows.Close()

	var groups []*iam_domain.GroupProxy
	for rows.Next() {
		group := &iam_domain.GroupProxy{}
		err := rows.Scan(&group.Id, &group.Name, &group.Type)
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
