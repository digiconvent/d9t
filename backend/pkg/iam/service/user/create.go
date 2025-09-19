package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamUserService) CreateUser(user *iam_domain.User) (*uuid.UUID, *core.Status) {
	id, status := s.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, status
	}
	return id, status
}
