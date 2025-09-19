package iam_user_service

import (
	"github.com/digiconvent/d9t/core"
	pagination "github.com/digiconvent/d9t/core/page"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (service *IamUserService) ListUsers(fs *iam_domain.UserFilterSort) (*pagination.Page[iam_domain.UserFacade], *core.Status) {
	page, status := service.repository.User.List(fs)
	if status.Err() {
		return nil, status
	}

	return page, status
}
