package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestPermissionGroupUpdate(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	permissionGroupID, _ := iamService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupUpdate",
		Abbr:        "PGU",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	status := iamService.PermissionGroup.UpdatePermissionGroup(permissionGroupID, &iam_domain.PermissionGroup{
		Name:        "PermissionGroupUpdate1",
		Abbr:        "PGUx",
		Description: "tset",
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	pgProfile, status := iamService.PermissionGroup.GetPermissionGroup(permissionGroupID)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if pgProfile == nil {
		t.Fatalf("No permission group found")
	}

	if pgProfile.Name != "PermissionGroupUpdate1" {
		t.Fatalf("Permission group name not updated")
	}

	if pgProfile.Abbr != "PGUx" {
		t.Fatalf("Permission group abbreviation not updated")
	}

	if pgProfile.Description != "tset" {
		t.Fatalf("Permission group description not updated")
	}
}
