package iam_domain

import "github.com/google/uuid"

type UserPassword struct {
	User     *uuid.UUID `json:"user"`
	Password string     `json:"password"`
}
