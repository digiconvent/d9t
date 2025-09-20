package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *userService) Create(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	return s.repo.Create(user)
}