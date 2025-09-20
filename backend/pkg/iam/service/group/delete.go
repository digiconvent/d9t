package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/google/uuid"
)

func (s *groupService) Delete(id *uuid.UUID) *core.Status {
	return s.repo.Delete(id)
}