package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestCreatePermissionGroup(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	testPermissionGroup := &iam_domain.PermissionGroup{
		Name:        "PermissionGroupCreate",
		Abbr:        "PG",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	}

	res, status := testService.PermissionGroup.CreatePermissionGroup(testPermissionGroup)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}
}
