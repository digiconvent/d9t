package iam_service

import (
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
)

type IamServices struct {
	User       iam_user_service.UserServiceInterface
	Group      iam_group_service.GroupServiceInterface
	Policy     iam_policy_service.PolicyServiceInterface
	Permission iam_permission_service.PermissionServiceInterface
}

func NewIamServices(repo *iam_repository.IamRepository) *IamServices {
	return &IamServices{
		User:       iam_user_service.NewUserService(repo.User),
		Group:      iam_group_service.NewGroupService(repo.Group),
		Policy:     iam_policy_service.NewPolicyService(repo.Policy),
		Permission: iam_permission_service.NewPermissionService(repo.Permission),
	}
}
