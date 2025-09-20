package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestDelete(t *testing.T) {
	service := tests.TestGroupService()

	group := &iam_domain.Group{
		Name:   "Delete Test Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}

	id, _ := service.Create(group)

	status := service.Delete(id)
	if !status.Ok() {
		t.Fatalf("Delete failed: %s", status.String())
	}

	_, status = service.Read(id)
	if status.Ok() {
		t.Fatal("Expected deleted group to not be found")
	}
}
