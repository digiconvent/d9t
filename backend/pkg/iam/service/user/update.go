package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *userService) Update(user *iam_domain.User) *core.Status {
	return s.repo.Update(user)
}

func (s *userService) SetEnabled(id *uuid.UUID, enabled bool) *core.Status {
	return s.repo.SetEnabled(id, enabled)
}