package iam_permission_service_test

import (
	"testing"

	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestListPermissions(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	permissions, status := iamService.Permission.List()

	if status.Err() {
		t.Errorf("ListPermissions() failed: %s", status.Message)
	}

	if len(permissions) == 0 {
		t.Errorf("ListPermissions() failed: no permissions found")
	}
}
