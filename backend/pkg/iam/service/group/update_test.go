package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestUpdate(t *testing.T) {
	service := tests.TestGroupService()

	group := &iam_domain.Group{
		Name:   "Update Test Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}

	id, _ := service.Create(group)

	updates := &iam_domain.Group{
		Id:     id,
		Name:   "Updated Group",
		Type:   "container",
		Parent: getGroupRoot(),
	}

	status := service.Update(updates)
	if !status.Ok() {
		t.Fatalf("Update failed: %s", status.String())
	}

	result, _ := service.Read(id)
	if result.Name != "Updated Group" {
		t.Errorf("Expected 'Updated Group', got %s", result.Name)
	}
}
