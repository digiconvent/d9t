package iam_repository

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/DigiConvent/testd9t/core/db"
	iam_credentials_repository "github.com/DigiConvent/testd9t/pkg/iam/repository/credentials"
	iam_permission_repository "github.com/DigiConvent/testd9t/pkg/iam/repository/permission"
	iam_permission_group_repository "github.com/DigiConvent/testd9t/pkg/iam/repository/permission_group"
	iam_policy_repository "github.com/DigiConvent/testd9t/pkg/iam/repository/policy"
	iam_user_repository "github.com/DigiConvent/testd9t/pkg/iam/repository/user"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
)

type IamRepository struct {
	db         db.DatabaseInterface
	privateKey *rsa.PrivateKey

	Credentials     iam_credentials_repository.IamCredentialsRepositoryInterface
	User            iam_user_repository.IamUserRepositoryInterface
	PermissionGroup iam_permission_group_repository.IamPermissionGroupRepositoryInterface
	Permission      iam_permission_repository.IamPermissionRepositoryInterface
	Policy          iam_policy_repository.IamPolicyRepositoryInterface
}

func NewIamRepository(db db.DatabaseInterface) *IamRepository {
	privateKeyPem, err := os.ReadFile(iam_setup.JwtPrivateKeyPath())
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(privateKeyPem)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return &IamRepository{
		db:              db,
		privateKey:      key,
		Credentials:     iam_credentials_repository.NewIamCredentialsRepository(db),
		Permission:      iam_permission_repository.NewIamPermissionRepository(db),
		PermissionGroup: iam_permission_group_repository.NewIamPermissionGroupRepository(db),
		Policy:          iam_policy_repository.NewIamPolicyRepository(db),
		User:            iam_user_repository.NewIamUserRepository(db),
	}
}
