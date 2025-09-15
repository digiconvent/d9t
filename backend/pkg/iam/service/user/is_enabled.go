package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IamUserService) IsEnabled(id *uuid.UUID) (bool, *core.Status) {
	enabled, status := service.repository.User.IsEnabled(id)
	return enabled, &status
}
