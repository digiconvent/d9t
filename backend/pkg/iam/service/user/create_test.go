package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

var testUser = &iam_domain.UserWrite{
	FirstName:    "FirstName",
	LastName:     "LastName",
	Emailaddress: "a@a.a",
}

func TestCreateUser(t *testing.T) {
	testService := iam_service_test.GetTestIamService()

	res, status := testService.User.Create(testUser)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}
}
