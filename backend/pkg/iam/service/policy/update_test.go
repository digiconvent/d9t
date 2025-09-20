package iam_policy_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_policy_service "github.com/digiconvent/d9t/pkg/iam/service/policy"
	"github.com/digiconvent/d9t/tests"
)

func TestUpdate(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_policy_service.NewPolicyService(repo.Policy)

	policy := &iam_domain.Policy{
		Name:          "Update Test Policy",
		VotesRequired: 1,
	}

	id, _ := service.Create(policy)

	updates := &iam_domain.Policy{
		Id:            id,
		Name:          "Updated Policy",
		VotesRequired: 2,
	}

	status := service.Update(updates)
	if !status.Ok() {
		t.Fatalf("Update failed: %s", status.String())
	}

	result, _ := service.Read(id)
	if result.Name != "Updated Policy" {
		t.Errorf("Expected 'Updated Policy', got %s", result.Name)
	}
}
