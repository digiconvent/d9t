package iam_user_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (service *IamUserService) ResetPassword(emailaddress string) (string, *core.Status) {
	user, status := service.repository.User.GetUserByEmailaddress(emailaddress)
	if status.Err() {
		return "", &status
	}

	token, status := service.repository.Credentials.ResetCredentials(&user.Id)

	return token, &status
}
