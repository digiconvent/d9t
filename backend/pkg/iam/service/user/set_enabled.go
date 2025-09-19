package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *IamUserService) SetEnabled(id *uuid.UUID, enabled bool) *core.Status {
	status := s.repository.User.SetEnabled(id, enabled)
	return status
}
