package iam_domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id           *uuid.UUID `json:"id"`
	Emailaddress string     `json:"emailaddress" validate:"required,email"`
	FirstName    string     `json:"first_name" validate:"required,alphaunicode"`
	LastName     string     `json:"last_name" validate:"required,alphaunicode"`
	Enabled      bool       `json:"enabled"`
}

type UserFacade struct {
	Id         uuid.UUID  `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	StatusId   *uuid.UUID `json:"status_id"`
	StatusName *string    `json:"status_name"`
	RoleId     *uuid.UUID `json:"role_id"`
	RoleName   *string    `json:"role_name"`
	Implied    bool       `json:"implied"`
}

type UserProfile struct {
	User       *User                    `json:"user"`
	UserStatus []*UserBecameStatusRead  `json:"status"`
	UserRole   []*UserHasRoleRead       `json:"role"`
	Groups     []*PermissionGroupFacade `json:"groups"`
	Policies   []*PolicyFacade          `json:"policies"`
}

type UserFilterSort struct {
	Filter struct {
		Emailaddress *string
		FirstName    *string
		LastName     *string
	}
	Sort struct {
		Field string
		Asc   bool
	}
	Page         int
	ItemsPerPage int
}
