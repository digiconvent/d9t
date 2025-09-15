package iam_auth_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestLoginUser(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	testUser, _ := testService.User.Create(&iam_domain.UserWrite{
		Emailaddress: "TestLoginUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	testService.Auth.SetUserPassword(testUser, "password123")

	uid, status := testService.Auth.LoginUser("TestLoginUser@test.test", "password123")

	if status.Err() {
		t.Fatal(status.Message)
	}

	if uid == nil {
		t.Fatal("Expected a result")
	}

	if uid.String() != testUser.String() {
		t.Fatal("Expected the same ID")
	}
}
