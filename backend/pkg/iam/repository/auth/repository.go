package iam_auth_repository

import (
	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/migrate_packages/db"
	uuid "github.com/google/uuid"
)

type IamAuthRepositoryInterface interface {
	ReadCredentials(emailaddress string) (*uuid.UUID, string, *core.Status)
	ResetCredentials(id *uuid.UUID) (string, *core.Status)
	SetCredentialEmailaddress(id *uuid.UUID, emailaddress string) *core.Status
	SetCredentialPassword(id *uuid.UUID, password string) *core.Status
	GetTelegramId(dataString, botToken string) (*int, *core.Status)
}

type IamAuthRepository struct {
	db db.DatabaseInterface
}

func NewIamAuthRepository(db db.DatabaseInterface) IamAuthRepositoryInterface {
	return &IamAuthRepository{
		db: db,
	}
}
