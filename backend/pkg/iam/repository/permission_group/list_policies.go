package iam_permission_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IamPermissionGroupRepository) ListPermissionGroupPolicies(arg *uuid.UUID) ([]*iam_domain.PolicyFacade, *core.Status) {
	var policies = make([]*iam_domain.PolicyFacade, 0)
	rows, err := r.db.Query(`select * from policy_has_permission_group where permission_group = ?`, arg.String())

	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var policy iam_domain.PolicyFacade
		err := rows.Scan(&policy.Name)
		if err != nil {
			return nil, core.InternalError(err.Error())
		}

		policies = append(policies, &policy)
	}

	return policies, core.StatusSuccess()
}
