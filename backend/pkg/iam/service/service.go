package iam_service

import (
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	iam_auth_service "github.com/digiconvent/d9t/pkg/iam/service/auth"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	iam_permission_group_service "github.com/digiconvent/d9t/pkg/iam/service/permission_group"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
)

type IamServices struct {
	Auth            iam_auth_service.IamAuthServiceInterface
	User            iam_user_service.IamUserServiceInterface
	PermissionGroup iam_permission_group_service.IamPermissionGroupServiceInterface
	Permission      iam_permission_service.IamPermissionServiceInterface
	Policy          iam_policy_service.IamPolicyServiceInterface
}

func NewIamServices(iamRepo *iam_repository.IamRepository) *IamServices {
	userServices := iam_user_service.NewUserService(iamRepo)
	authServices := iam_auth_service.NewAuthService(iamRepo)
	permissionServices := iam_permission_service.NewIamPermissionService(iamRepo)
	permissionGroupServices := iam_permission_group_service.NewIamPermissionGroupService(iamRepo)
	policyServices := iam_policy_service.NewPolicyService(iamRepo)
	return &IamServices{
		Auth:            authServices,
		PermissionGroup: permissionGroupServices,
		Permission:      permissionServices,
		Policy:          policyServices,
		User:            userServices,
	}
}
