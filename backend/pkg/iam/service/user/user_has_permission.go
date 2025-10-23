package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *userService) UserHasPermission(user *uuid.UUID, permission string) ([]*iam_domain.PolicyProxy, *core.Status) {
	return s.repo.ReadPoliciesWithPermission(user, permission)
}