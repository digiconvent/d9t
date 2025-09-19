package iam_user_service_test

import (
	"testing"

	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	iam_service_test_utils "github.com/digiconvent/d9t/pkg/iam/service/test"
	"github.com/google/uuid"
)

func TestRead(t *testing.T) {
	iamService := iam_service_test_utils.GetTestIamService()

	fakeUser := iam_domain.User{
		Emailaddress: "TestGetUser@test.test",
		FirstName:    "Test",
		LastName:     "GetUser",
	}

	id, _ := iamService.User.Create(&fakeUser)

	user, status := iamService.User.Read(id)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if user == nil {
		t.Fatal("User is nil")
	}

	if *user.Id != *id {
		t.Errorf("ID is not equal")
	}

	if user.Emailaddress != fakeUser.Emailaddress {
		t.Errorf("Email is not equal")
	}

	if user.FirstName != fakeUser.FirstName {
		t.Errorf("FirstName is not equal")
	}

	if user.LastName != fakeUser.LastName {
		t.Errorf("LastName is not equal")
	}

	// test fake uuid

	id = nil
	user, status = iamService.User.Read(id)

	if !status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if user != nil {
		t.Fatal("User is not nil")
	}

	unknownId, _ := uuid.NewV7()

	user, status = iamService.User.Read(&unknownId)

	if !status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if status.Code != 404 {
		t.Errorf("Status code is not 404")
	}

	if user != nil {
		t.Fatal("User is not nil")
	}

	// also test admin user
	adminId := uuid.Nil
	admin, status := iamService.User.Read(&adminId)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if admin == nil {
		t.Fatal("User is nil")
	}

	if admin.Emailaddress != "" {
		t.Errorf("Email is not empty")
	}

	profile, status := iamService.User.GetUserProfile(&adminId)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if profile == nil {
		t.Fatal("Profile is nil")
	}

	if profile.User == nil {
		t.Fatal("User is nil")
	}
}
