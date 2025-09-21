package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type Policy struct {
	Id            *uuid.UUID `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	VotesRequired int        `json:"votes_required"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type PolicyProxy struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

type PolicyProfile struct {
	Policy      *Policy       `json:"policy"`
	Groups      []*GroupProxy `json:"groups"`
	Permissions []string      `json:"permissions"`
}
