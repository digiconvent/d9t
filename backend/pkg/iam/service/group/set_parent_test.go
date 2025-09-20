package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestSetParent(t *testing.T) {
	service := tests.TestGroupService()

	parent := &iam_domain.Group{Name: "Parent Group", Type: "container", Parent: getGroupRoot()}
	parentId, _ := service.Create(parent)

	child := &iam_domain.Group{Name: "Child Group", Type: "role", Parent: getGroupRoot()}
	childId, _ := service.Create(child)

	status := service.SetParent(childId, parentId)
	if !status.Ok() {
		t.Fatalf("SetParent failed: %s", status.String())
	}
}
