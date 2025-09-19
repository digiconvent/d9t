package iam_permission_group_service_test

import (
	"os"
	"path"

	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	iam_permission_group_service "github.com/digiconvent/d9t/pkg/iam/service/permission_group"
	"github.com/digiconvent/migrate_packages/db"
)

func NewTestIamPermissionGroupService() iam_permission_group_service.IamPermissionGroupServiceInterface {
	testDb, _ := db.New(path.Join(os.TempDir(), "digiconvent_test", "iam", "iam.db"))
	iamAuthRepository := iam_repository.NewIamRepository(testDb)
	return iam_permission_group_service.NewIamPermissionGroupService(iamAuthRepository)
}
