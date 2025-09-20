package iam_policy_repository

import (
	"database/sql"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *policyRepository) Read(id *uuid.UUID) (*iam_domain.Policy, *core.Status) {
	query := `select id, name, description, votes_required from policies where id = ?`

	policy := &iam_domain.Policy{}
	err := r.db.QueryRow(query, id.String()).Scan(
		&policy.Id,
		&policy.Name,
		&policy.Description,
		&policy.VotesRequired,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NotFoundError("policy not found")
		}
		return nil, core.InternalError("failed to read policy")
	}

	return policy, core.StatusSuccess()
}

func (r *policyRepository) ReadProxies() ([]*iam_domain.PolicyProxy, *core.Status) {
	query := `select id, name from policies`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, core.InternalError("failed to read policies")
	}
	defer rows.Close()

	var policies []*iam_domain.PolicyProxy
	for rows.Next() {
		policy := &iam_domain.PolicyProxy{}
		err := rows.Scan(&policy.Id, &policy.Name)
		if err != nil {
			return nil, core.InternalError("failed to scan policy")
		}
		policies = append(policies, policy)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate policies")
	}

	return policies, core.StatusSuccess()
}
