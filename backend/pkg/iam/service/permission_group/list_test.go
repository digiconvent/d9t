package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test "github.com/digiconvent/d9t/pkg/iam/service/test"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestListPermissionGroups(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupList1",
		Abbr:        "PG",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})
	testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupList2",
		Abbr:        "PG",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	permissionGroups, status := testService.PermissionGroup.ListPermissionGroups()

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if len(permissionGroups) == 0 {
		t.Fatalf("No permission groups found")
	}

	if len(permissionGroups) < 2 {
		t.Fatalf("Expected at least 2 permission groups, instead got %v", len(permissionGroups))
	}

	found1, found2 := false, false
	for _, permissionGroup := range permissionGroups {
		if permissionGroup.Name == "PermissionGroupList1" {
			found1 = true
		}
		if permissionGroup.Name == "PermissionGroupList2" {
			found2 = true
		}
	}

	if !found1 {
		t.Fatalf("Expected PermissionGroupList1, not found")
	}

	if !found2 {
		t.Fatalf("Expected PermissionGroupList2, not found")
	}
}
