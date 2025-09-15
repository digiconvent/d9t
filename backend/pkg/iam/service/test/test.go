package iam_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
	"github.com/google/uuid"

	setup_pkg "github.com/DigiConvent/testd9t/setup"
)

var testDB db.DatabaseInterface

func getRootPermissionGroup() string {
	testService := GetTestIamService()
	facades, _ := testService.PermissionGroup.ListPermissionGroups()
	for _, facade := range facades {
		if facade.Name == "root" {
			return facade.Id.String()
		}
	}

	// testRepo := getTestRepo("iam")
	// id, _ := testRepo.PermissionGroup.Create(&iam_domain.PermissionGroup{
	// 	Name: "root",
	// })

	return "" // id.String()
}

func GetRootPermissionGroupUuid() *uuid.UUID {
	root := getRootPermissionGroup()
	res := uuid.MustParse(root)
	return &res
}

func GetTestIamService() *iam_service.IamServices {
	os.Setenv(constants.CERTIFICATES_PATH, "/tmp/testd9t/certificates")
	if testDB == nil {
		iam_setup.Setup()
		testDB, _ = setup_pkg.SetupTestDatabase("iam")
	}
	repo := iam_repository.NewIamRepository(testDB)
	s := iam_service.NewIamServices(repo)
	return &s
}

func TestMain(m *testing.M) {
	GetTestIamService()
	defer testDB.DeleteDatabase()
	m.Run()
}
