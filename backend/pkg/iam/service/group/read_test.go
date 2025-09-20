package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestRead(t *testing.T) {
	service := tests.TestGroupService()

	group := &iam_domain.Group{
		Name:   "Read Test Group",
		Type:   "role",
		Parent: getGroupRoot(),
	}

	id, _ := service.Create(group)

	result, status := service.Read(id)
	if !status.Ok() {
		t.Fatalf("Read failed: %s", status.String())
	}

	if result.Name != group.Name {
		t.Errorf("Expected %s, got %s", group.Name, result.Name)
	}
}
