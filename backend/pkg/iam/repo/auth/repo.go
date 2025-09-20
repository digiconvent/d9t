package iam_auth_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/migrate_packages/db"
	"github.com/google/uuid"
)

type AuthRepositoryInterface interface {
	SetPassword(userPassword *iam_domain.UserPassword) *core.Status
	GetPasswordHash(user *uuid.UUID) (string, *core.Status)
}

type authRepository struct {
	db db.DatabaseInterface
}

func NewAuthRepository(database db.DatabaseInterface) AuthRepositoryInterface {
	return &authRepository{db: database}
}