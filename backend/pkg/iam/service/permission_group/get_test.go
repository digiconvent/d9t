package iam_permission_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test "github.com/digiconvent/d9t/pkg/iam/service/test"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestGetPermissionGroup(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	res, status := testService.PermissionGroup.CreatePermissionGroup(&iam_domain.PermissionGroup{
		Name:        "PermissionGroupGet",
		Abbr:        "PG",
		Description: "test",
		Parent:      iam_service_test.GetRootPermissionGroupUuid(),
	})

	if res == nil {
		t.Log(status.Message)
		t.Fatal("Expected a result, instead got", status.Message)
	}

	permissionGroup, status := testService.PermissionGroup.GetPermissionGroup(res)

	if status.Err() {
		t.Fatal("Unable to get permission group:", status.Message)
	}

	if permissionGroup.Name != "PermissionGroupGet" {
		t.Fatal("Expected PermissionGroupGet, instead I got ", permissionGroup.Name)
	}

	if permissionGroup.Abbr != "PG" {
		t.Fatal("Expected PG, instead I got ", permissionGroup.Abbr)
	}

	if permissionGroup.Description != "test" {
		t.Fatal("Expected test, instead I got ", permissionGroup.Description)
	}

	randomFailingId, _ := uuid.NewV7()
	permissionGroup, status = testService.PermissionGroup.GetPermissionGroup(&randomFailingId)

	if !status.Err() {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	if status.Code != 404 {
		t.Fatal("Expected 404")
	}

	if permissionGroup != nil {
		t.Fatal("Expected nil, instead I got ", permissionGroup)
	}
}
