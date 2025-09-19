package iam_domain

import "github.com/google/uuid"

type Policy struct {
	Id            *uuid.UUID `json:"id"`
	Name          string     `json:"name" validate:"alphaunicode,required"`
	Description   string     `json:"description"`
	RequiredVotes int        `json:"required_votes" validate:"numeric,required,min=-100,max=100"`
}

type PolicyFacade struct {
	Id            *uuid.UUID `json:"id"`
	Name          string     `json:"name"`
	RequiredVotes int        `json:"required_votes"`
}
