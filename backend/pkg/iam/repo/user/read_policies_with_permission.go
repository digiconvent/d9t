package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) ReadPoliciesWithPermission(user *uuid.UUID, permission string) ([]*iam_domain.PolicyProxy, *core.Status) {
	query := `
		select distinct p.id, p.name
		from policies p
		join group_has_policy ghp on p.id = ghp.policy
		join group_has_user ghu on ghp."group" = ghu."group"
		join policy_has_permission php on p.id = php.policy
		where ghu."user" = ? and php.permission = ?`

	rows, err := r.db.Query(query, user.String(), permission)
	if err != nil {
		return nil, core.InternalError("failed to read user policies with permission: " + err.Error())
	}
	defer rows.Close()

	var policies []*iam_domain.PolicyProxy
	for rows.Next() {
		policy := &iam_domain.PolicyProxy{}
		err := rows.Scan(&policy.Id, &policy.Name)
		if err != nil {
			return nil, core.InternalError("failed to scan user policy with permission")
		}
		policies = append(policies, policy)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate user policies with permission")
	}

	return policies, core.StatusSuccess()
}
