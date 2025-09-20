package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestCreate(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_user_service.NewUserService(repo.User)

	user := &iam_domain.User{
		Email:     "create@example.com",
		FirstName: "Create",
		LastName:  "Test",
		Enabled:   true,
	}

	id, status := service.Create(user)
	if !status.Ok() {
		t.Fatalf("Create failed: %s", status.String())
	}

	if id == nil {
		t.Fatal("Expected user ID to be returned")
	}
}
