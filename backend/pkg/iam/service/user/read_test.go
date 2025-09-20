package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/digiconvent/d9t/tests"
)

func TestRead(t *testing.T) {
	service := tests.TestUserService()

	user := &iam_domain.User{
		Email:     "read@example.com",
		FirstName: "Read",
		LastName:  "Test",
		Enabled:   true,
	}

	id, _ := service.Create(user)

	result, status := service.Read(id)
	if !status.Ok() {
		t.Fatalf("Read failed: %s", status.String())
	}

	if result.Email != user.Email {
		t.Errorf("Expected %s, got %s", user.Email, result.Email)
	}
}
