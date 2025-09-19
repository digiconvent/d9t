package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *IamPolicyRepository) List() ([]*iam_domain.Policy, *core.Status) {
	var policies []*iam_domain.Policy
	rows, err := r.db.Query(`select id, name, description, required_votes from policies`)

	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		policy := &iam_domain.Policy{}
		err = rows.Scan(&policy.Id, &policy.Name, &policy.Description, &policy.RequiredVotes)
		if err != nil {
			return nil, core.InternalError(err.Error())
		}
		policies = append(policies, policy)
	}
	return policies, core.StatusSuccess()
}
