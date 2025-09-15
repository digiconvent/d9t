package iam_auth_service

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s *IamAuthService) GenerateJwt(userId *uuid.UUID) (string, *core.Status) {
	if userId == nil {
		return "", core.UnprocessableContentError("ID is required")
	}
	privKey := s.repository.GetPrivateKey()

	user, _ := s.repository.User.GetUserByID(userId)

	if user == nil || !user.Enabled {
		return "", core.UnauthorizedError("User is not enabled")
	}

	telegramId, _ := s.repository.User.GetUserTelegramID(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"id":   userId.String(),
		"user": user,
		"tgid": telegramId,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString(privKey)
	if err != nil {
		return "", core.BadRequestError(err.Error())
	}

	return tokenString, core.StatusSuccess()
}
