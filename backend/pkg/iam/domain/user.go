package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        *uuid.UUID `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Telegram  *int       `json:"telegram"`
	Enabled   bool       `json:"enabled"`
	JoinedAt  time.Time  `json:"joined_at"`
}

type UserProxy struct {
	Id        *uuid.UUID `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
}

type UserProfile struct {
	User        *User         `json:"user"`
	Groups      []*GroupProxy `json:"groups"`
	Permissions []string      `json:"permissions"`
}
