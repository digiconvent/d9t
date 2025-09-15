package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (s *IamUserService) CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		return nil, &core.Status{Code: 422, Message: err.Error()}
	}

	id, status := s.repository.User.CreateUser(user)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}
	return id, &status
}
