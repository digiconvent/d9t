package iam_policy_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (i *IamPolicyService) Delete(id *uuid.UUID) *core.Status {
	return i.repository.Policy.Delete(id)
}
