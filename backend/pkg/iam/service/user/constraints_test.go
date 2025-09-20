package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestDuplicateEmailConstraint(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_user_service.NewUserService(repo.User)

	user1 := &iam_domain.User{
		Email:     "duplicate@example.com",
		FirstName: "First",
		LastName:  "User",
		Enabled:   true,
	}

	_, status := service.Create(user1)
	if !status.Ok() {
		t.Fatalf("First user creation failed: %s", status.String())
	}

	user2 := &iam_domain.User{
		Email:     "duplicate@example.com",
		FirstName: "Second",
		LastName:  "User",
		Enabled:   true,
	}

	_, status = service.Create(user2)
	if status.Ok() {
		t.Fatal("Expected duplicate email creation to fail")
	}

	if status.Code != 409 && status.Code != 500 {
		t.Errorf("Expected conflict (409) or internal error (500), got %d", status.Code)
	}
}
