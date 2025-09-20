package iam_permission_service

import (
	"github.com/digiconvent/d9t/core"
)

func (s *permissionService) Delete(permission string) *core.Status {
	return s.repo.Delete(permission)
}