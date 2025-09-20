package iam_group_repository

import (
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *groupRepository) ReadUsers(id *uuid.UUID) ([]*iam_domain.UserProxy, *core.Status) {
	query := `select u.id, u.email, u.first_name, u.last_name from users u join group_has_user ghu on u.id = ghu."user" where ghu."group" = ?`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return nil, core.InternalError("failed to read group users: " + err.Error())
	}
	defer rows.Close()

	var users []*iam_domain.UserProxy
	for rows.Next() {
		user := &iam_domain.UserProxy{}
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, core.InternalError("failed to scan group user")
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate group users")
	}

	return users, core.StatusSuccess()
}
