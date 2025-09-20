package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestDelete(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	service := iam_user_service.NewUserService(repo.User)

	user := &iam_domain.User{
		Email:     "delete@example.com",
		FirstName: "Delete",
		LastName:  "Test",
		Enabled:   true,
	}

	id, _ := service.Create(user)

	status := service.Delete(id)
	if !status.Ok() {
		t.Fatalf("Delete failed: %s", status.String())
	}

	_, status = service.Read(id)
	if status.Ok() {
		t.Fatal("Expected deleted user to not be found")
	}
}
