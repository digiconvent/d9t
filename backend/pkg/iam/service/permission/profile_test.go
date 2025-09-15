package iam_permission_service_test

import (
	"testing"

	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestGetPermissionProfile(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	permission, status := iamService.Permission.Profile("idam")
	if !status.Err() || permission != nil {
		t.Fatalf("GetPermissionProfile() succeeded where it should have failed")
	}

	permission, status = iamService.Permission.Profile("iam")

	if status.Err() {
		t.Fatalf("GetPermissionProfile() failed: %s", status.Message)
	}

	if permission == nil || permission.Permission == nil {
		t.Fatalf("GetPermissionProfile() failed: no permission found")
	}

	if permission.Permission.Name != "iam" {
		t.Fatalf("GetPermissionProfile() failed: wrong permission returned")
	}

	if permission.Descendants == nil {
		t.Fatalf("GetPermissionProfile() failed: no descendants found")
	}

	if len(permission.Descendants) == 0 {
		t.Fatalf("GetPermissionProfile() failed: no descendants found")
	}
}
