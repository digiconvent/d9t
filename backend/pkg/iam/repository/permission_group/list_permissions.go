package iam_permission_group_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) ListPermissionGroupPolicies(arg *uuid.UUID) ([]*iam_domain.Policy, core.Status) {
	var policies = make([]*iam_domain.Policy, 0)
	rows, err := r.db.Query(`select * from policy_has_policy_group where policy_group = ? order by implied desc`, arg.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var policy iam_domain.Policy
		// err := rows.Scan(&policy.Name, &policy.Implied)
		// if err != nil {
		// 	return nil, *core.InternalError(err.Error())
		// }

		policies = append(policies, &policy)
	}

	return policies, *core.StatusSuccess()
}
