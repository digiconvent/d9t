package iam_user_handles

import (
	context "github.com/digiconvent/d9t/api/context"
)

type ReadRequest struct {
}

func Read(ctx *context.Context) {
	user, status := ctx.Services.Iam.User.Read(ctx.Id)
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	ctx.Json(200, user)
}
