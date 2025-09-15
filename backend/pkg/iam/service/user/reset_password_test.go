package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestResetPassword(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	uid, status := iamService.User.Create(&iam_domain.UserWrite{
		Emailaddress: "TestResetPassword@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	_, status = iamService.User.Read(uid)
	if status.Err() {
		t.Fatal(status.Message)
	}

	token, status := iamService.Auth.ResetPassword("TestResetPassword@test.test")
	if status.Err() {
		t.Fatal(status.Message)
	}

	if token == "" {
		t.Fatal("Expected a token")
	}
}
