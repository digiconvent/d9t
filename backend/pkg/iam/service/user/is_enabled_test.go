package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
)

func TestIsEnabled(t *testing.T) {
	iamService := iam_service_test_utils.GetTestIamService()

	user := &iam_domain.User{
		Emailaddress: "UserIsEnabled@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.User.Create(user)

	userRead, _ := iamService.User.Read(userId)
	if userRead.Enabled {
		t.Errorf("User should initially be disabled")
	}

	isEnabled, _ := iamService.User.IsEnabled(userId)
	if isEnabled {
		t.Errorf("User should be disabled initially")
	}

	status := iamService.User.SetEnabled(userId, true)
	if status.Err() {
		t.Errorf("Error enabling user: %v", status)
	}

	isEnabled, _ = iamService.User.IsEnabled(userId)
	if !isEnabled {
		t.Errorf("User should be enabled")
	}

	status = iamService.User.SetEnabled(userId, false)
	if status.Err() {
		t.Errorf("Error disabling user: %v", status)
	}

	isEnabled, _ = iamService.User.IsEnabled(userId)
	if isEnabled {
		t.Errorf("User should be disabled")
	}

	status = iamService.User.SetEnabled(nil, true)
	if !status.Err() {
		t.Errorf("Expected error when setting enabled for nil user")
	}

	isEnabled, status = iamService.User.IsEnabled(nil)
	if !status.Err() {
		t.Errorf("Expected error when getting enabled for nil user")
	}
	if isEnabled {
		t.Errorf("Expected false when getting enabled for nil user")
	}
}
