package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamUserService) UpdateUser(id *uuid.UUID, user *iam_domain.User) *core.Status {

	status := s.repository.User.UpdateUser(id, user)
	return status
}
