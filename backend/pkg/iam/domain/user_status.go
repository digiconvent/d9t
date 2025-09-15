package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type UserStatusProfile struct {
	*PermissionGroupProfile
	UserStatus *UserStatusRead         `json:"user_role"`
	History    []*UserBecameStatusRead `json:"history"`
}

type UserStatusRead struct {
	PermissionGroup
	Archived bool `json:"archived"`
}

type UserStatusWrite struct {
	PermissionGroup
	Archived bool `json:"archived"`
}

type UserBecameStatusRead struct {
	UserStatus uuid.UUID  `json:"user_status"`
	User       UserFacade `json:"user"`
	Comment    string     `json:"comment"`
	Start      time.Time  `json:"start"`
	End        *time.Time `json:"end"` // if this is nil, the user is currently that status
}

type UserBecameStatusWrite struct {
	UserStatus uuid.UUID  `json:"user_status"`
	User       uuid.UUID  `json:"user"`
	Comment    string     `json:"comment"`
	Start      time.Time  `json:"start"`
	End        *time.Time `json:"end"`
}
