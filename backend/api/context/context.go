package context

import (
	"net/http"

	"github.com/digiconvent/d9t/core"
	"github.com/digiconvent/d9t/meta/services"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/google/uuid"
)

type Context struct {
	Author   *iam_domain.User
	Id       *uuid.UUID
	Request  *http.Request
	Response http.ResponseWriter
	Services *services.Services
}

func (ctx *Context) Respond(response string, status *core.Status) {
	ctx.Response.Write([]byte(response))
	ctx.Response.WriteHeader(status.Code)
}

func (ctx *Context) HandleStatus(status *core.Status) {
	ctx.Response.WriteHeader(status.Code)
	ctx.Response.Write([]byte(status.Message))
}
