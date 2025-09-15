package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamUserService) UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) *core.Status {
	if user == nil {
		return core.UnprocessableContentError("iam.user.update.missing_user")
	}
	if user.Emailaddress == "" {
		return core.UnprocessableContentError("iam.user.update.missing_email")
	}
	if !core_utils.ValidEmail(user.Emailaddress) {
		return core.UnprocessableContentError("iam.user.update.invalid_email")
	}

	status := s.repository.User.UpdateUser(id, user)
	return &status
}
