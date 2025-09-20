package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *userService) RemoveGroup(user, group *uuid.UUID) *core.Status {
	return s.repo.RemoveGroup(user, group)
}