package iam_permission_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	"github.com/digiconvent/d9t/tests"
)

func TestList(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_permission_service.NewPermissionService(repo.Permission)

	permission := &iam_domain.Permission{
		Permission: "list.test.permission",
	}

	service.Create(permission)

	permissions, status := service.List()
	if !status.Ok() {
		t.Fatalf("List failed: %s", status.String())
	}

	found := false
	for _, p := range permissions {
		if p.Permission == "list.test.permission" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Permission not found in list")
	}
}
