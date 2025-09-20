package iam_permission_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	"github.com/digiconvent/d9t/tests"
)

func TestDelete(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_permission_service.NewPermissionService(repo.Permission)

	permission := &iam_domain.Permission{
		Permission: "delete.test.permission",
	}

	service.Create(permission)

	status := service.Delete(permission.Permission)
	if !status.Ok() {
		t.Fatalf("Delete failed: %s", status.String())
	}
}
