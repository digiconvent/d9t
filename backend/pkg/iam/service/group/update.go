package iam_group_service

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
)

func (s *groupService) Update(group *iam_domain.Group) *core.Status {
	return s.repo.Update(group)
}