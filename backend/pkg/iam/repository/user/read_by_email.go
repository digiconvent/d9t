package iam_user_repository

import (
	"strings"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func (r *IamUserRepository) GetUserByEmailaddress(emailaddress string) (*iam_domain.User, *core.Status) {
	var user = &iam_domain.User{}
	row := r.db.QueryRow(`select id, first_name, last_name, enabled from users where emailaddress = ?`, strings.ToLower(emailaddress))

	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Enabled,
	)
	if err != nil {
		return nil, core.NotFoundError("No user found with email: " + emailaddress)
	}

	user.Emailaddress = emailaddress

	return user, core.StatusSuccess()
}
