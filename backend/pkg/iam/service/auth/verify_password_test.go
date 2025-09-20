package iam_auth_service_test

import (
	"testing"

	iam_auth_service "github.com/digiconvent/d9t/pkg/iam/service/auth"
)

func TestVerifyPassword(t *testing.T) {
	service := iam_auth_service.NewAuthService()

	password := "my_secure_password"
	wrongPassword := "wrong_password"

	hash, status := service.HashPassword(password)
	if !status.Ok() {
		t.Fatalf("Failed to hash password: %s", status.String())
	}

	status = service.VerifyPassword(password, hash)
	if !status.Ok() {
		t.Fatalf("Expected password verification to succeed, got: %s", status.String())
	}

	status = service.VerifyPassword(wrongPassword, hash)
	if status.Ok() {
		t.Fatal("Expected password verification to fail with wrong password")
	}

	if status.Code != 401 {
		t.Errorf("Expected 401 Unauthorized, got: %d", status.Code)
	}
}
