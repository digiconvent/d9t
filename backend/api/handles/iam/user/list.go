package iam_user_handles

import (
	context "github.com/digiconvent/d9t/api/context"
)

type ListRequest struct {
}

func List(ctx *context.Context) {
	users, status := ctx.Services.Iam.User.ReadProxies()
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	ctx.Json(200, users)
}
