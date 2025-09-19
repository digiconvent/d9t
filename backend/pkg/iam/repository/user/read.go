package iam_user_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IamUserRepository) GetUserByID(id *uuid.UUID) (*iam_domain.User, *core.Status) {
	if id == nil {
		return nil, core.UnprocessableContentError("ID is required")
	}
	var user = &iam_domain.User{}
	row := r.db.QueryRow(`select id, emailaddress, first_name, last_name, enabled from users where id = ?`, id.String())

	err := row.Scan(
		&user.Id,
		&user.Emailaddress,
		&user.FirstName,
		&user.LastName,
		&user.Enabled,
	)
	if err != nil {
		return nil, core.NotFoundError("User not found with that id")
	}

	return user, core.StatusSuccess()
}
