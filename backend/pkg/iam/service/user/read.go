package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *userService) Read(id *uuid.UUID) (*iam_domain.User, *core.Status) {
	return s.repo.Read(id)
}

func (s *userService) ReadProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status) {
	profile := &iam_domain.UserProfile{}
	return profile, core.StatusSuccess()
}

func (s *userService) ReadProxies() ([]*iam_domain.UserProxy, *core.Status) {
	return s.repo.ReadProxies()
}

func (s *userService) ReadByEmail(email string) (*iam_domain.User, *core.Status) {
	return s.repo.ReadByEmail(email)
}
