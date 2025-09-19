package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *IamUserRepository) AddStatusToUser(d *iam_domain.UserBecameStatusWrite) *core.Status {
	_, err := r.db.Exec(`insert into permission_group_has_user ("user", permission_group, start, comment) values (?, ?, ?, ?)`, d.User, d.UserStatus, d.Start, d.Comment)

	if err != nil {
		return core.InternalError(err.Error())
	}

	return core.StatusSuccess()
}
