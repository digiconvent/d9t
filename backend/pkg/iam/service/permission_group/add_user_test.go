package iam_permission_group_service_test

import (
	"fmt"
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestAddUserToPermissionGroup(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	pgData := &iam_domain.PermissionGroup{
		Name:        "PermissionGroupAddUser",
		Abbr:        "PGA",
		Description: "test",
		Parent:      iam_service_test_utils.GetRootPermissionGroupUuid(),
	}

	pg, status := testService.PermissionGroup.CreatePermissionGroup(pgData)

	if pg == nil {
		t.Fatal("Expected a result, instead got ", status)
	}

	user, _ := testService.User.Create(&iam_domain.User{
		Emailaddress: "PermissionGroupAddUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	if user == nil {
		t.Fatal("Expected a result")
	}

	testService.PermissionGroup.AddUserToPermissionGroup(pg, user)

	id := iam_service_test_utils.GetRootPermissionGroupUuid().String()
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
	rootId := uuid.MustParse(iam_service_test_utils.GetRootPermissionGroupUuid().String())
	rootProfile, _ := testService.PermissionGroup.GetPermissionGroupProfile(&rootId)

	if rootProfile == nil {
		t.Fatal("Expected a result")
	}

	// we inserted Test McTest but Admin McAdmin is also a user in the root group
	if len(rootProfile.Users) != 2 {
		for _, u := range rootProfile.Users {
			fmt.Println(u.FirstName, u.LastName)
		}
		t.Fatal("Expected 1 user, instead I got", rootProfile.Users)
	}
}
