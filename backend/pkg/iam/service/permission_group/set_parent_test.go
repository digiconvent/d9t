package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestSetPermissionGroupParent(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	permissionGroupGrandParentID, _ := iamService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupGrandParent",
		Abbr:        "PGP",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})
	permissionGroupChildID, _ := iamService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupParentSetChild",
		Abbr:        "PGP",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	permissionGroupParentID, _ := iamService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupParent",
		Abbr:        "PGP",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	status := iamService.PermissionGroup.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		Id:     permissionGroupChildID,
		Parent: permissionGroupParentID,
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	status = iamService.PermissionGroup.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		Id:     permissionGroupParentID,
		Parent: permissionGroupGrandParentID,
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	pgProfile, status := iamService.PermissionGroup.GetPermissionGroupProfile(permissionGroupChildID)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if pgProfile == nil {
		t.Fatalf("No permission group found")
	}

	if len(pgProfile.Ancestors) != 4 {
		t.Fatalf("Expected 4 permission group, instead got %v", len(pgProfile.Ancestors))
	}
}
