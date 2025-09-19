package iam_repository

import (
	"crypto/rsa"

	"github.com/digiconvent/d9t/meta/environment"
	iam_auth_repository "github.com/digiconvent/d9t/pkg/iam/repository/auth"
	iam_permission_repository "github.com/digiconvent/d9t/pkg/iam/repository/permission"
	iam_permission_group_repository "github.com/digiconvent/d9t/pkg/iam/repository/permission_group"
	iam_policy_repository "github.com/digiconvent/d9t/pkg/iam/repository/policy"
	iam_user_repository "github.com/digiconvent/d9t/pkg/iam/repository/user"
	"github.com/digiconvent/d9t/utils/sec"
	"github.com/digiconvent/migrate_packages/db"
)

type IamRepository struct {
	db         db.DatabaseInterface
	privateKey *rsa.PrivateKey

	Auth            iam_auth_repository.IamAuthRepositoryInterface
	User            iam_user_repository.IamUserRepositoryInterface
	PermissionGroup iam_permission_group_repository.IamPermissionGroupRepositoryInterface
	Permission      iam_permission_repository.IamPermissionRepositoryInterface
	Policy          iam_policy_repository.IamPolicyRepositoryInterface
}

func NewIamRepository(db db.DatabaseInterface) *IamRepository {
	return &IamRepository{
		db:              db,
		privateKey:      sec.StringToPrivateKey(environment.Env.JwtPk), // this is from the environment variables
		Auth:            iam_auth_repository.NewIamAuthRepository(db),
		Permission:      iam_permission_repository.NewIamPermissionRepository(db),
		PermissionGroup: iam_permission_group_repository.NewIamPermissionGroupRepository(db),
		Policy:          iam_policy_repository.NewIamPolicyRepository(db),
		User:            iam_user_repository.NewIamUserRepository(db),
	}
}
