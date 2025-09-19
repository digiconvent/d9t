package iam_auth_service

import (
	"github.com/digiconvent/d9t/core"
)

func (service *IamAuthService) ResetPassword(emailaddress string) (string, *core.Status) {
	user, status := service.repository.User.GetUserByEmailaddress(emailaddress)
	if status.Err() {
		return "", status
	}

	token, status := service.repository.Auth.ResetCredentials(user.Id)

	return token, status
}
