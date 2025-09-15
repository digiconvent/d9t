package iam_permission_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestCreatePermission(t *testing.T) {
	service := iam_service_test.GetTestIamService()
	permissionName := "service.permission.create_test"

	permissions, _ := service.Permission.List()

	for _, permission := range permissions {
		if permission.Name == permissionName {
			t.Errorf("Expected permission %v to be not existent", permissionName)
		}
	}

	status := service.Permission.Create(&iam_domain.PermissionWrite{
		Name:        permissionName,
		Description: "test",
		Meta:        "",
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	permissions, _ = service.Permission.List()

	expectToExist := []string{"service.permission.create_test", "service.permission", "service"}

	for permission := range expectToExist {
		if !contains(permissions, expectToExist[permission]) {
			for _, p := range permissions {
				t.Logf("Permission: %v", p.Name)
			}
			t.Fatalf("Expected permission %v to exist", expectToExist[permission])
		}
	}
}

func contains(s []*iam_domain.PermissionRead, e string) bool {
	for _, a := range s {
		if a.Name == e {
			return true
		}
	}
	return false
}
