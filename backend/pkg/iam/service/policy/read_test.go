package iam_policy_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	"github.com/digiconvent/d9t/tests"
)

func TestRead(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_policy_service.NewPolicyService(repo.Policy)

	policy := &iam_domain.Policy{
		Name:          "Read Test Policy",
		VotesRequired: 1,
	}

	id, _ := service.Create(policy)

	result, status := service.Read(id)
	if !status.Ok() {
		t.Fatalf("Read failed: %s", status.String())
	}

	if result.Name != policy.Name {
		t.Errorf("Expected %s, got %s", policy.Name, result.Name)
	}
}
