package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) ReadPolicies(id *uuid.UUID) ([]*iam_domain.PolicyProxy, *core.Status) {
	query := `select p.id, p.name from policies p join group_has_policy ghp on p.id = ghp."policy" where ghp."group" = ?`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return nil, core.InternalError("failed to read group policies")
	}
	defer rows.Close()

	var policies []*iam_domain.PolicyProxy
	for rows.Next() {
		policy := &iam_domain.PolicyProxy{}
		err := rows.Scan(&policy.Id, &policy.Name)
		if err != nil {
			return nil, core.InternalError("failed to scan group policy")
		}
		policies = append(policies, policy)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate group policies")
	}

	return policies, core.StatusSuccess()
}
