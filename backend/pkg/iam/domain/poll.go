package iam_domain

import "github.com/google/uuid"

type Poll struct {
	Id     *uuid.UUID `json:"id"`
	Policy *uuid.UUID `json:"policy"`
}
