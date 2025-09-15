package iam_user_service

import (
	"github.com/google/uuid"
)

func (service *IamUserService) UserHasPermission(id *uuid.UUID, permission string) bool {
	return service.repository.User.UserHasPermission(id, permission)
}
