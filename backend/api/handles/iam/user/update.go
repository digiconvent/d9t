package iam_user_handles

import (
	context "github.com/digiconvent/d9t/api/context"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Update(ctx *context.Context) {
	data, status := context.ParseAndValidate[UpdateRequest](ctx.Request)
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	id := ctx.Id
	status = ctx.Services.Iam.User.Update(&iam_domain.User{
		Id:        id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	})

	ctx.HandleStatus(status)
}
