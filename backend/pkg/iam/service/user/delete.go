package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *userService) Delete(id *uuid.UUID) *core.Status {
	return s.repo.Delete(id)
}