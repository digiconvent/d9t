package iam_auth_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestVerifyJwt(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	_, status := testService.Auth.VerifyJwt("badtoken")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	_, status = testService.Auth.VerifyJwt("")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	testUser := &iam_domain.User{
		Emailaddress: "VerifyJwt@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := testService.User.Create(testUser)
	testService.User.SetEnabled(id, true)
	token, _ := testService.Auth.GenerateJwt(id)
	if !status.Err() {
		t.Fatal("this should fail because the user is not enabled")
	}

	theId, status := testService.Auth.VerifyJwt(token)
	if status.Err() {
		t.Fatal(status.Message)
	}

	if theId == nil {
		t.Fatal("Expected a result")
	}

	if theId.String() != id.String() {
		t.Fatal("Expected ", id.String(), " instead I got ", theId.String())
	}
}
