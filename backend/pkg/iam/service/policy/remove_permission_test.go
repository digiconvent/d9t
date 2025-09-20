package iam_policy_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_permission_service "github.com/digiconvent/d9t/pkg/iam/service/permission"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	"github.com/digiconvent/d9t/tests"
)

func TestRemovePermission(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	policyService := iam_policy_service.NewPolicyService(repo.Policy)
	permissionService := iam_permission_service.NewPermissionService(repo.Permission)

	policy := &iam_domain.Policy{Name: "Remove Permission Policy", VotesRequired: 1}
	policyId, _ := policyService.Create(policy)

	permission := &iam_domain.Permission{Permission: "remove.permission.test"}
	permissionService.Create(permission)

	policyService.AddPermission(policyId, permission.Permission)

	status := policyService.RemovePermission(policyId, permission.Permission)
	if !status.Ok() {
		t.Fatalf("RemovePermission failed: %s", status.String())
	}
}
