package iam_policy_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestCreate(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()
	policy := &iam_domain.Policy{
		Name:          "TestCreatePolicy",
		RequiredVotes: -200,
	}

	id, err := testService.Policy.Create(policy)
	if err == nil {
		t.Fatalf("expected an error")
	}
	if id != nil {
		t.Fatalf("expected id to be nil")
	}
	policy.RequiredVotes = -100
	id, err = testService.Policy.Create(policy)
	if err.Err() {
		t.Fatalf("did not expect an error, instead got %v", err.Code)
	}

	if id == nil {
		t.Fatal("did not expect id to be nil")
	}
}
