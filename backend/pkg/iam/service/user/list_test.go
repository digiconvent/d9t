package iam_user_service_test

import (
	"testing"

	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestListUsers(t *testing.T) {
	iamService := iam_service_test_utils.GetTestIamService()

	userList, status := iamService.User.ListUsers(nil)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if userList == nil {
		t.Fatal("Expected a page of users, got nil")
	}

}
