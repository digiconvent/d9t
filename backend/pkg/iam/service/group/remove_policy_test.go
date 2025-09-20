package iam_group_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_group_service "github.com/digiconvent/d9t/pkg/iam/service/group"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	"github.com/digiconvent/d9t/tests"
)

func TestRemovePolicy(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	groupService := iam_group_service.NewGroupService(repo.Group)
	policyService := iam_policy_service.NewPolicyService(repo.Policy)

	group := &iam_domain.Group{Name: "Remove Policy Group", Type: "role", Parent: getGroupRoot()}
	groupId, _ := groupService.Create(group)

	policy := &iam_domain.Policy{Name: "Remove Policy Test", VotesRequired: 1}
	policyId, _ := policyService.Create(policy)

	groupService.AddPolicy(groupId, policyId)

	status := groupService.RemovePolicy(groupId, policyId)
	if !status.Ok() {
		t.Fatalf("RemovePolicy failed: %s", status.String())
	}
}
