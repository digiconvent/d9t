package iam_policy_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	"github.com/digiconvent/d9t/tests"
)

func TestAddPermission(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	policyService := iam_policy_service.NewPolicyService(repo.Policy)
	permissionService := iam_permission_service.NewPermissionService(repo.Permission)

	policy := &iam_domain.Policy{Name: "Add Permission Policy", VotesRequired: 1}
	policyId, _ := policyService.Create(policy)

	permission := &iam_domain.Permission{Permission: "add.permission.test"}
	permissionService.Create(permission)

	status := policyService.AddPermission(policyId, permission.Permission)
	if !status.Ok() {
		t.Fatalf("AddPermission failed: %s", status.String())
	}
}
