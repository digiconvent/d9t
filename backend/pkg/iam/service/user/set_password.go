package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (service *IamUserService) SetUserPassword(id *uuid.UUID, rawPassword string) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("iam.user.set_password.missing_user")
	}
	if rawPassword == "" {
		return core.UnprocessableContentError("iam.user.set_password.empty")
	}

	password, err := hashedPassword(rawPassword)
	if err != nil {
		return core.InternalError(err.Error())
	}

	status := service.repository.Credentials.SetCredentialPassword(id, password)
	if status.Err() {
		return &status
	}

	return core.StatusNoContent()
}
