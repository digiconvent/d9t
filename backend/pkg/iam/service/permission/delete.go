package iam_permission_service

import (
	"github.com/digiconvent/d9t/core"
)

func (service *IamPermissionService) Delete(name string) *core.Status {
	status := service.repository.Permission.DeletePermission(name)
	return status
}
