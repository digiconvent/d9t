package iam_user_handles

import (
	"github.com/digiconvent/d9t/api/context"
	"github.com/digiconvent/d9t/core"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

type CreateRequest struct {
	Email string `json:"email" validate:"email"`
}

func Create(ctx *context.Context) {
	data, status := context.ParseAndValidate[CreateRequest](ctx.Request)
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	id, status := ctx.Services.Iam.User.Create(&iam_domain.User{
		Email: data.Email,
	})
	if status.Err() {
		_, check := ctx.Services.Iam.User.ReadByEmail(data.Email)
		if !check.Err() {
			ctx.HandleStatus(core.ConflictError("iam.user.email.duplicate"))
			return
		}

		ctx.HandleStatus(status)
		return
	}
	ctx.Respond(status.Code, id.String())
}
