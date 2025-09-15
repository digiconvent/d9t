package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestSetEnabled(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	user := &iam_domain.UserWrite{
		Emailaddress: "SetUserEnabled@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.User.Create(user)

	userRead, _ := iamService.User.Read(userId)
	if userRead.Enabled {
		t.Errorf("User should initially be disabled")
	}

	iamService.User.IsEnabled(userId)

	status := iamService.User.SetEnabled(userId, true)
	if status.Err() {
		t.Errorf("Error enabling user: %v", status)
	}

	userRead, _ = iamService.User.Read(userId)
	if !userRead.Enabled {
		t.Errorf("User should be enabled, instead got: %v", userRead.Enabled)
	}

	status = iamService.User.SetEnabled(userId, false)
	if status.Err() {
		t.Errorf("Error disabling user: %v", status)
	}

	userRead, _ = iamService.User.Read(userId)
	if userRead.Enabled {
		t.Errorf("User should be disabled")
	}

	status = iamService.User.SetEnabled(nil, true)
	if !status.Err() {
		t.Errorf("Expected error when setting enabled for nil user")
	}
}
