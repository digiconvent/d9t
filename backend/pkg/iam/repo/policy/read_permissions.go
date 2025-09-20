package iam_policy_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (r *policyRepository) ReadPermissions(id *uuid.UUID) ([]string, *core.Status) {
	permissions := []string{}
	return permissions, nil
}
