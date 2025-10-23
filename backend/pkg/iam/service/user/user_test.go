package iam_user_service_test

import (
	"fmt"
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_repository "github.com/digiconvent/d9t/pkg/iam/repo"
	iam_user_service "github.com/digiconvent/d9t/pkg/iam/service/user"
	"github.com/digiconvent/d9t/tests"
)

func TestUserCrud(t *testing.T) {
	db := tests.GetTestDatabase("iam")
	repo := iam_repository.NewIamRepository(db)
	userService := iam_user_service.NewUserService(repo.User)

	t.Run("Create", func(t *testing.T) {
		user := &iam_domain.User{
			Email:     "test@example.com",
			FirstName: "Test",
			LastName:  "User",
			Enabled:   true,
		}

		id, status := userService.Create(user)
		if !status.Ok() {
			t.Fatalf("Failed to create user: %s", status.String())
		}

		if id == nil {
			t.Fatal("Expected user ID to be returned")
		}

		result, status := userService.Read(id)
		if !status.Ok() {
			t.Fatalf("Failed to read created user: %s", status.String())
		}

		if result.Email != user.Email {
			t.Errorf("Expected email %s, got %s", user.Email, result.Email)
		}
	})

	t.Run("ReadNonExistent", func(t *testing.T) {
		user := &iam_domain.User{
			Email:     "read-non-existent@example.com",
			FirstName: "Read",
			LastName:  "User",
			Enabled:   true,
		}

		id, status := userService.Create(user)
		fmt.Println(status)

		_, status = userService.Read(id)
		if status.Err() {
			t.Fatalf("Failed to read existing user: %s", status.String())
		}

		fakeId := *id
		fakeId[0] = ^fakeId[0]

		_, status = userService.Read(&fakeId)
		if status.Ok() {
			t.Fatal("Expected non-existent user read to fail")
		}

		if status.Code != 404 {
			t.Errorf("Expected not found status 404, got %d", status.Code)
		}
	})

	t.Run("Update", func(t *testing.T) {
		user := &iam_domain.User{
			Email:     "crud-update@example.com",
			FirstName: "Update",
			LastName:  "User",
			Enabled:   true,
		}

		id, _ := userService.Create(user)

		updates := &iam_domain.User{
			Id:        id,
			Email:     user.Email,
			FirstName: "Updated",
			LastName:  "Name",
			Enabled:   user.Enabled,
		}

		status := userService.Update(updates)
		if !status.Ok() {
			t.Fatalf("Failed to update user: %s", status.String())
		}

		result, _ := userService.Read(id)
		if result.FirstName != "Updated" {
			t.Errorf("Expected first name 'Updated', got %s", result.FirstName)
		}
		if result.LastName != "Name" {
			t.Errorf("Expected last name 'Name', got %s", result.LastName)
		}
		if result.Email != user.Email {
			t.Errorf("Expected email %s, got %s", user.Email, result.Email)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		user := &iam_domain.User{
			Email:     "crud-delete@example.com",
			FirstName: "Delete",
			LastName:  "User",
			Enabled:   true,
		}

		id, _ := userService.Create(user)

		status := userService.Delete(id)
		if !status.Ok() {
			t.Fatalf("Failed to delete user: %s", status.String())
		}

		_, status = userService.Read(id)
		if status.Ok() {
			t.Fatal("Expected deleted user to not be found")
		}
	})
}

func TestUserGroupMembership(t *testing.T) {
	t.Skip("Group membership tests need complex time-based logic - skipping for now")
}
