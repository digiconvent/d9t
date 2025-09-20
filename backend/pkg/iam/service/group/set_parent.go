package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *groupService) SetParent(group, parent *uuid.UUID) *core.Status {
	return s.repo.SetParent(group, parent)
}