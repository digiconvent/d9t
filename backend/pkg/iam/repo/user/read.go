package iam_user_repository

import (
	"database/sql"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *userRepository) Read(id *uuid.UUID) (*iam_domain.User, *core.Status) {
	query := `select id, email, first_name, last_name, password_hash, telegram, enabled, joined_at from users where id = ?`

	user := &iam_domain.User{}
	err := r.db.QueryRow(query, id.String()).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
		&user.Telegram,
		&user.Enabled,
		&user.JoinedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NotFoundError("user not found")
		}
		return nil, core.InternalError("failed to read user")
	}

	return user, core.StatusSuccess()
}

func (r *userRepository) ReadByEmail(email string) (*iam_domain.User, *core.Status) {
	query := `select id, email, first_name, last_name, password_hash, telegram, enabled, joined_at from users where lower(email) = lower(?)`

	user := &iam_domain.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
		&user.Telegram,
		&user.Enabled,
		&user.JoinedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, core.NotFoundError("user not found")
		}
		return nil, core.InternalError("failed to read user")
	}

	return user, core.StatusSuccess()
}

func (r *userRepository) ReadProxies() ([]*iam_domain.UserProxy, *core.Status) {
	query := `select id, email, first_name, last_name from users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, core.InternalError("failed to read users")
	}
	defer rows.Close()

	var users []*iam_domain.UserProxy
	for rows.Next() {
		user := &iam_domain.UserProxy{}
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, core.InternalError("failed to scan user")
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, core.InternalError("failed to iterate users")
	}

	return users, core.StatusSuccess()
}
