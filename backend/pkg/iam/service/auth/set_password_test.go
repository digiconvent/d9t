package iam_auth_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestSetUserPassword(t *testing.T) {
	iamService := iam_service_test_utils.GetTestIamService()

	testUser, _ := iamService.User.Create(&iam_domain.User{
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
