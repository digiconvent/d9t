package iam_user_service

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IamUserService) GetUserProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status) {
	if id == nil {
		return nil, &core.Status{Code: 422, Message: "ID is required"}
	}

	userRead, status := s.repository.User.GetUserByID(id)
	if status.Err() {
		return nil, &status
	}
	fmt.Println("This is not working")

	userPermissions, status := s.repository.User.ListUserPermissions(id)
	if status.Err() {
		return nil, &status
	}

	userGroups, status := s.repository.User.ListUserGroups(id)
	if status.Err() {
		return nil, &status
	}

	return &iam_domain.UserProfile{
		User:        userRead,
		Groups:      userGroups,
		Permissions: userPermissions,
	}, core.StatusSuccess()
}
