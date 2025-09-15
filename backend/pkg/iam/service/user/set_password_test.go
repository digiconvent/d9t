package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestSetUserPassword(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	testUser, _ := iamService.User.Create(&iam_domain.UserWrite{
		Emailaddress: "TestSetUserPassword@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	status := iamService.Auth.SetUserPassword(testUser, "password123")

	if status.Err() {
		t.Fatal(status.Message)
	}

	status = iamService.Auth.SetUserPassword(nil, "password123")

	if !status.Err() {
		t.Fatal("Expected an error")
	}
}
