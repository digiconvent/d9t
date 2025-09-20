package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestAddPolicy(t *testing.T) {
	groupService := tests.TestGroupService()
	policyService := tests.TestPolicyService()

	group := &iam_domain.Group{Name: "Add Policy Group", Type: "role", Parent: getGroupRoot()}
	groupId, status := groupService.Create(group)
	if status.Err() {
		t.Fatal(status.Message)
	}
	t.Log(groupId)
	policy := &iam_domain.Policy{Name: "Add Policy Test", VotesRequired: 1}
	policyId, status := policyService.Create(policy)
	if status.Err() {
		t.Fatal(status.Message)
	}

	status = groupService.AddPolicy(groupId, policyId)
	if !status.Ok() {
		t.Fatalf("AddPolicy failed: %s", status.String())
	}
}
