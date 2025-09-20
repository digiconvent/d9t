package iam_group_service_test

import (
	"testing"

	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestGroupServiceReadProfile(t *testing.T) {
	var status *core.Status
	groupService := tests.TestGroupService()
	userService := tests.TestUserService()
	policyService := tests.TestPolicyService()

	root := getGroupRoot()
	parent := &iam_domain.Group{Name: "Parent", Type: "container", Parent: root}
	parentId, _ := groupService.Create(parent)
	group := &iam_domain.Group{Name: "Group", Type: "container", Parent: parentId}
	groupId, _ := groupService.Create(group)
	child := &iam_domain.Group{Name: "Child", Type: "role", Parent: groupId}
	childId, _ := groupService.Create(child)

	status = groupService.SetParent(groupId, parentId)
	if status.Err() {
		t.Fatal(status.Message)
	}
	status = groupService.SetParent(childId, groupId)
	if status.Err() {
		t.Fatal(status.Message)
	}
	user := &iam_domain.User{Email: "service@example.com", FirstName: "Service", LastName: "User", Enabled: true}
	userId, _ := userService.Create(user)
	status = groupService.AddUser(childId, userId)
	if status.Err() {
		t.Fatal(status.Message)
	}

	policy := &iam_domain.Policy{Name: "Policy", VotesRequired: 1}
	policyId, _ := policyService.Create(policy)
	groupService.AddPolicy(groupId, policyId)

	profile, status := groupService.ReadProfile(groupId)
	if !status.Ok() {
		t.Fatalf("ReadProfile failed: %s", status.String())
	}

	if profile.Group == nil {
		t.Fatal("Expected Group to be populated in service")
	}

	if profile.Group.Name != "Group" {
		t.Errorf("Expected group name 'Group', got %s", profile.Group.Name)
	}

	if len(profile.Ascendants) != 2 {
		t.Errorf("Expected 2 ascendant, got %d", len(profile.Ascendants))
	}

	if profile.Ascendants[0].Name != "root" {
		t.Errorf("Expected ascendant 'Parent', got %s", profile.Ascendants[0].Name)
	}

	if len(profile.Descendants) != 1 {
		t.Errorf("Expected 1 descendant, got %d", len(profile.Descendants))
	}

	if profile.Descendants[0].Name != "Child" {
		t.Errorf("Expected descendant 'Child', got %s", profile.Descendants[0].Name)
	}
}
