package context

import (
	"encoding/json"
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

func (ctx *Context) Respond(code int, body string) {
	ctx.Response.WriteHeader(code)
	ctx.Response.Write([]byte(body))
}

func (ctx *Context) Json(code int, a any) {
	body, err := json.Marshal(a)
	if err != nil {
		ctx.Respond(400, "could not jsonify this")
		return
	}
	ctx.Response.WriteHeader(code)
	ctx.Response.Write([]byte(body))
}

func (ctx *Context) HandleStatus(status *core.Status) {
	ctx.Response.WriteHeader(status.Code)
	ctx.Response.Write([]byte(status.Message))
}
