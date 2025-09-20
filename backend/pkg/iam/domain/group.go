package iam_domain

import "github.com/google/uuid"

type Group struct {
	Id          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Parent      *uuid.UUID `json:"parent"`
	Description *string    `json:"description"`
}

type GroupProxy struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Type string     `json:"type"`
}

type GroupProfile struct {
	Group       *Group         `json:"group"`
	Ascendants  []*GroupProxy  `json:"ascendants"`
	Descendants []*GroupProxy  `json:"descendants"`
	Users       []*UserProxy   `json:"users"`
	Policies    []*PolicyProxy `json:"policies"`
}
