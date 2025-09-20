package iam_permission_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
)

func (s *permissionService) Create(permission *iam_domain.Permission) *core.Status {
	return s.repo.Create(permission)
}