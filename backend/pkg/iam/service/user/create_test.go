package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

var testUser = &iam_domain.User{
	FirstName:    "FirstName",
	LastName:     "LastName",
	Emailaddress: "a@a.a",
}

func TestCreateUser(t *testing.T) {
	testService := iam_service_test_utils.GetTestIamService()

	res, status := testService.User.Create(testUser)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}
}
