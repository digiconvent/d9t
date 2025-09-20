package iam_repository

import (
	iam_auth_repository "github.com/digiconvent/d9t/pkg/iam/repo/auth"
	iam_group_repository "github.com/digiconvent/d9t/pkg/iam/repo/group"
	iam_permission_repository "github.com/digiconvent/d9t/pkg/iam/repo/permission"
	iam_policy_repository "github.com/digiconvent/d9t/pkg/iam/repo/policy"
	iam_user_repository "github.com/digiconvent/d9t/pkg/iam/repo/user"
	"github.com/digiconvent/migrate_packages/db"
)

type IamRepository struct {
	User       iam_user_repository.UserRepositoryInterface
	Group      iam_group_repository.GroupRepositoryInterface
	Policy     iam_policy_repository.PolicyRepositoryInterface
	Permission iam_permission_repository.PermissionRepositoryInterface
	Auth       iam_auth_repository.AuthRepositoryInterface
}

func NewIamRepository(database db.DatabaseInterface) *IamRepository {
	return &IamRepository{
		User:       iam_user_repository.NewUserRepository(database),
		Group:      iam_group_repository.NewGroupRepository(database),
		Policy:     iam_policy_repository.NewPolicyRepository(database),
		Permission: iam_permission_repository.NewPermissionRepository(database),
		Auth:       iam_auth_repository.NewAuthRepository(database),
	}
}
