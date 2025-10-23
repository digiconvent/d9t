package tests

import (
	iam_group_repository "github.com/digiconvent/d9t/pkg/iam/repo/group"
	iam_permission_repository "github.com/digiconvent/d9t/pkg/iam/repo/permission"
	iam_policy_repository "github.com/digiconvent/d9t/pkg/iam/repo/policy"
	iam_user_repository "github.com/digiconvent/d9t/pkg/iam/repo/user"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
)

var iamDb = GetTestDatabase("iam")

func TestUserService() iam_user_service.UserServiceInterface {
	repo := iam_user_repository.NewUserRepository(iamDb)
	userService := iam_user_service.NewUserService(repo)
	return userService
}

func TestGroupService() iam_group_service.GroupServiceInterface {
	repo := iam_group_repository.NewGroupRepository(iamDb)
	groupService := iam_group_service.NewGroupService(repo)
	return groupService
}

func TestPolicyService() iam_policy_service.PolicyServiceInterface {
	repo := iam_policy_repository.NewPolicyRepository(iamDb)
	policyService := iam_policy_service.NewPolicyService(repo)
	return policyService
}

func TestPermissionService() iam_permission_service.PermissionServiceInterface {
	repo := iam_permission_repository.NewPermissionRepository(iamDb)
	permissionService := iam_permission_service.NewPermissionService(repo)
	return permissionService
}
