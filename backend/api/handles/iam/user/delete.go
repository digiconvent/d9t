package iam_user_handles

import (
	context "github.com/digiconvent/d9t/api/context"
)

func Delete(ctx *context.Context) {
	status := ctx.Services.Iam.User.Delete(ctx.Id)
	ctx.Respond(status.Code, status.Message)
}
