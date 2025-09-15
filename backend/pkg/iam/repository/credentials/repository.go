package iam_credentials_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	uuid "github.com/google/uuid"
)

type IamCredentialsRepositoryInterface interface {
	ReadCredentials(emailaddress string) (*uuid.UUID, string, core.Status)
	ResetCredentials(id *uuid.UUID) (string, core.Status)
	SetCredentialEmailaddress(id *uuid.UUID, emailaddress string) core.Status
	SetCredentialPassword(id *uuid.UUID, password string) core.Status
}

type IamCredentialsRepository struct {
	db db.DatabaseInterface
}

func NewIamCredentialsRepository(db db.DatabaseInterface) IamCredentialsRepositoryInterface {
	return &IamCredentialsRepository{
		db: db,
	}
}
