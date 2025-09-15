package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestAddUserToPermissionGroup(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	pg, _ := testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupAddUser",
		Abbr:        "PGA",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	if pg == nil {
		t.Fatal("Expected a result")
	}

	user, _ := testService.User.Create(&iam_domain.UserWrite{
		Emailaddress: "PermissionGroupAddUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	if user == nil {
		t.Fatal("Expected a result")
	}

	testService.PermissionGroup.AddUserToPermissionGroup(pg, user)

	id := iam_service_test.GetRootPermissionGroupUuid().String()
	parsedId, _ := uuid.Parse(id)

	userStatus, status := testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupAddUserTest",
		Abbr:        "PGAUT",
		Description: "testxs",
		Parent:      &parsedId,
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	if userStatus == nil {
		t.Fatal("Expected a result")
	}

	status = testService.PermissionGroup.AddUserToPermissionGroup(userStatus, user)

	if status.Err() {
		t.Fatal(status.Message)
	}

	// get profile and count users
	rootId := uuid.MustParse(iam_service_test.GetRootPermissionGroupUuid().String())
	rootProfile, _ := testService.PermissionGroup.GetPermissionGroupProfile(&rootId)

	if rootProfile == nil {
		t.Fatal("Expected a result")
	}

	if len(rootProfile.Users) != 1 {
		t.Fatal("Expected 1 user, instead I got ", len(rootProfile.Users))
	}
}
