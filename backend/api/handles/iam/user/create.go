package iam_user_handles

import (
	"github.com/digiconvent/d9t/api/context"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
)

func Create(ctx *context.Context) {
	data, status := context.ParseAndValidate[iam_domain.User](ctx.Request)
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	id, status := ctx.Services.Iam.User.Create(data)
	if status.Err() {
		ctx.HandleStatus(status)
		return
	}
	ctx.Respond(id.String(), status)
}
func Read(ctx *context.Context) {
	ctx.Response.Write([]byte("Reading user: " + ctx.Id.String()))
}
func Update(ctx *context.Context) {
	ctx.Response.Write([]byte("Updating user: " + ctx.Id.String()))
}
func Delete(ctx *context.Context) {
	ctx.Response.Write([]byte("creating user"))
}
func List(ctx *context.Context) {
	ctx.Response.Write([]byte("Listing users"))
}
func AddRole(ctx *context.Context) {
	ctx.Response.Write([]byte("creating user"))
}
