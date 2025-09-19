package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestGetUserProfile(t *testing.T) {
	iamService := iam_service_test_utils.GetTestIamService()

	user := &iam_domain.User{
		Emailaddress: "GetUserProfile@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.User.Create(user)

	profile, err := iamService.User.GetUserProfile(userId)
	if err.Err() {
		t.Fatalf("GetUserProfile failed: %v", err)
	}

	if profile == nil {
		t.Fatalf("GetUserProfile failed: profile is nil")
	}

	if profile.User == nil {
		t.Fatalf("GetUserProfile failed: user is nil")
	}
}
