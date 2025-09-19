package iam_permission_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestDeletePermission(t *testing.T) {
	service := iam_service_test_utils.GetTestIamService()

	permissionName := "some.test.permission"
	service.Permission.Create(&iam_domain.PermissionWrite{
		Name:        permissionName,
		Description: "test",
		Meta:        "",
	})

	status := service.Permission.Delete(permissionName)
	if status != nil && status.Err() {
		t.Fatal("Error deleting permission: " + status.Message)
	}

	if status.Code != 204 {
		t.Fatal("Expected 204, got ", status.Code)
	}

	permissions, _ := service.Permission.List()

	for _, permission := range permissions {
		if permission.Name == permissionName {
			t.Fatal("Expected permission", permissionName, "to be deleted")
		}
	}
}
