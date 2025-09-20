package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *groupService) Create(group *iam_domain.Group) (*uuid.UUID, *core.Status) {
	return s.repo.Create(group)
}