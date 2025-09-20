package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestCreate(t *testing.T) {
	service := tests.TestGroupService()

	group := &iam_domain.Group{
		Name:   "Create Test Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}

	id, status := service.Create(group)
	if !status.Ok() {
		t.Fatalf("Create failed: %s", status.String())
	}

	if id == nil {
		t.Fatal("Expected group ID to be returned")
	}
}
