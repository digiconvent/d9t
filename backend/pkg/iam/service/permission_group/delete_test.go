package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test "github.com/digiconvent/d9t/pkg/iam/service/test"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestDeletePermissionGroup(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	res, _ := testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupDelete",
		Abbr:        "PG",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
		Policies:    []*uuid.UUID{},
	})

	if res == nil {
		t.Fatal("Expected a result")
	}

	status := testService.PermissionGroup.DeletePermissionGroup(res)

	if status.Err() {
		t.Fatal(status.Message)
	}

	_, status = testService.PermissionGroup.GetPermissionGroup(res)

	if !status.Err() {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	if status.Code != 404 {
		t.Fatal("Expected 404")
	}
}
