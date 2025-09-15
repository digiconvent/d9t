package iam_auth_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestGenerateJwt(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	testUser := &iam_domain.UserWrite{
		Emailaddress: "GenerateJwt@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := testService.User.Create(testUser)

	if id == nil {
		t.Fatal("Unable to create a user")
	}

	// this will fail if the user is not enabled
	_, status := testService.Auth.GenerateJwt(id)

	if status == nil {
		t.Fatal("status should not be nil")
	}

	testService.User.SetEnabled(id, true)

	token, _ := testService.Auth.GenerateJwt(id)

	if token == "" {
		t.Fatal("Expected a token")
	}

	_, status = testService.Auth.GenerateJwt(nil)
	if !status.Err() {
		t.Fatal("Expected an error")
	}
}
