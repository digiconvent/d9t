package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type PermissionGroup struct {
	Id          *uuid.UUID   `json:"id"`
	Name        string       `json:"name"`
	Abbr        string       `json:"abbr"`
	Description string       `json:"description"`
	Parent      *uuid.UUID   `json:"parent"`
	Policies    []*uuid.UUID `json:"permissions"`
	Meta        string       `json:"meta"`
}

type PermissionGroupSetParent struct {
	Id     *uuid.UUID `json:"id"`
	Parent *uuid.UUID `json:"parent"`
}

type PermissionGroupFacade struct {
	Id      uuid.UUID  `json:"id"`
	Name    string     `json:"name"`
	Abbr    string     `json:"abbr"`
	Meta    *string    `json:"meta"`
	Implied bool       `json:"implied"`
	Parent  *uuid.UUID `json:"parent"`
}

type PermissionGroupProfile struct {
	PermissionGroup *PermissionGroup         `json:"permission_group"`
	Users           []*UserFacade            `json:"users"`
	Ancestors       []*PermissionGroupFacade `json:"ancestors"`
	Descendants     []*PermissionGroupFacade `json:"descendants"`
	Policies        []*PolicyFacade          `json:"policies"`
}

type AddUserToPermissionGroupWrite struct {
	PermissionGroup *uuid.UUID `json:"permission_group"`
	User            *uuid.UUID `json:"user"`
	Start           *time.Time `json:"start"`
	End             *time.Time `json:"end"`
}
