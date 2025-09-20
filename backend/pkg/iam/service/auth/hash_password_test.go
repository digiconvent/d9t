package iam_auth_service_test

import (
	"testing"

	iam_auth_service "github.com/digiconvent/d9t/pkg/iam/service/auth"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	service := iam_auth_service.NewAuthService()

	password := "my_secure_password"

	hash, status := service.HashPassword(password)

	if !status.Ok() {
		t.Fatalf("Expected success, got: %s", status.String())
	}

	if hash == "" {
		t.Fatal("Expected hash to be returned")
	}

	if hash == password {
		t.Fatal("Hash should not equal original password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Fatalf("Hash verification failed: %v", err)
	}
}
