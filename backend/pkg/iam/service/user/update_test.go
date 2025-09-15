package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_service_test "github.com/DigiConvent/testd9t/pkg/iam/service/test"
)

func TestUpdateUser(t *testing.T) {
	iamService := iam_service_test.GetTestIamService()

	user := &iam_domain.UserWrite{
		Emailaddress: "TestUpdateUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.User.Create(user)

	user.Emailaddress = "TestUpdateUser@test.test1"
	user.FirstName = "Updated"
	// user.LastName = "McUpdated2"

	status := iamService.User.UpdateUser(userId, user)
	if status.Err() {
		t.Fatal("Error updating user", status.Message)
	}

	updatedUser, _ := iamService.User.Read(userId)
	if updatedUser.FirstName != user.FirstName ||
		updatedUser.LastName != user.LastName ||
		updatedUser.Emailaddress != user.Emailaddress {
		t.Fatal("User not updated")
	}

	status = iamService.User.UpdateUser(nil, user)
	if !status.Err() {
		t.Fatal("Should have errored")
	}

	status = iamService.User.UpdateUser(userId, nil)
	if !status.Err() {
		t.Fatal("Should have errored")
	}
}
