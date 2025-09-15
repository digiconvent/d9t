package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestGetPermissionGroupProfile(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	profileId, _ := testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupProfile",
		Abbr:        "PG",
		Description: "test",
		Policies:    []*uuid.UUID{},
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	permissionGroupProfile, status := testService.PermissionGroup.GetPermissionGroupProfile(profileId)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if permissionGroupProfile == nil {
		t.Fatalf("No permission group found")
	}

	if permissionGroupProfile.PermissionGroup.Name != "PermissionGroupProfile" {
		t.Fatalf("Expected PermissionGroupProfile, instead got %v", permissionGroupProfile.PermissionGroup.Name)
	}

	if permissionGroupProfile.PermissionGroup.Abbr != "PG" {
		t.Fatalf("Expected PG, instead got %v", permissionGroupProfile.PermissionGroup.Abbr)
	}

	if permissionGroupProfile.PermissionGroup.Description != "test" {
		t.Fatalf("Expected test, instead got %v", permissionGroupProfile.PermissionGroup.Description)
	}

	if permissionGroupProfile.PermissionGroup.Id != profileId {
		t.Fatalf("Expected %v, instead got %v", profileId, permissionGroupProfile.PermissionGroup.Id)
	}

	if len(permissionGroupProfile.Permissions) != 3 {
		t.Fatalf("Expected 3, instead got %v", len(permissionGroupProfile.Permissions))
	}
}
