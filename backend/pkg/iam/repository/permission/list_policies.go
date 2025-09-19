package iam_permission_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *IamPermissionRepository) ListPolicies(name string) ([]*iam_domain.PolicyFacade, *core.Status) {
	rows, err := r.db.Query(`select id, name, required_votes from policies where permission = ?`, name)

	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	defer rows.Close()

	policies := make([]*iam_domain.PolicyFacade, 0)
	for rows.Next() {
		policy := &iam_domain.PolicyFacade{}
		err := rows.Scan(&policy.Id, &policy.Name, &policy.RequiredVotes)
		if err != nil {
			return nil, core.InternalError(err.Error())
		}
		policies = append(policies, policy)
	}
	return policies, core.StatusSuccess()
}
