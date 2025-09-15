package iam_auth_service

import (
	"errors"

	"github.com/DigiConvent/testd9t/core"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s *IamAuthService) VerifyJwt(token string) (*uuid.UUID, *core.Status) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return &s.repository.GetPrivateKey().PublicKey, nil
	})

	if err != nil {
		return nil, core.UnauthorizedError(err.Error())
	}

	id := parsedToken.Claims.(jwt.MapClaims)["id"]
	if id == nil {
		return nil, core.UnauthorizedError("invalid token")
	}
	userId := uuid.MustParse(id.(string))

	return &userId, core.StatusSuccess()
}
