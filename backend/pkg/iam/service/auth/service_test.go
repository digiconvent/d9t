package iam_auth_service_test

import (
	"os"
	"path"

	iam_repository "github.com/digiconvent/d9t/pkg/iam/repository"
	iam_auth_service "github.com/digiconvent/d9t/pkg/iam/service/auth"
	"github.com/digiconvent/migrate_packages/db"
)

func NewTestIamAuthService() iam_auth_service.IamAuthServiceInterface {
	testDb, _ := db.New(path.Join(os.TempDir(), "digiconvent_test", "iam", "iam.db"))
	iamAuthRepository := iam_repository.NewIamRepository(testDb)
	return iam_auth_service.NewAuthService(iamAuthRepository)
}
