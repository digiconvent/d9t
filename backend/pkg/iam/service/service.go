package iam_service

import (
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_auth_service "github.com/DigiConvent/testd9t/pkg/iam/service/auth"
	iam_permission_service "github.com/DigiConvent/testd9t/pkg/iam/service/permission"
	iam_permission_group_service "github.com/DigiConvent/testd9t/pkg/iam/service/permission_group"
	iam_policy_service "github.com/DigiConvent/testd9t/pkg/iam/service/policy"
	iam_user_service "github.com/DigiConvent/testd9t/pkg/iam/service/user"
)

type IamServices struct {
	Auth            iam_auth_service.IamAuthServiceInterface
	User            iam_user_service.IamUserServiceInterface
	PermissionGroup iam_permission_group_service.IamPermissionGroupServiceInterface
	Permission      iam_permission_service.IamPermissionServiceInterface
	Policy          iam_policy_service.IamPolicyServiceInterface
}

func NewIamServices(iamRepo *iam_repository.IamRepository) IamServices {
	userServices := iam_user_service.NewUserService(iamRepo)
	permissionGroupServices := iam_permission_group_service.NewIamPermissionGroupService(iamRepo)
	return IamServices{
		User:            userServices,
		PermissionGroup: permissionGroupServices,
	}
}
