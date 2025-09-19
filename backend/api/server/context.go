package server

import (
	"io"
	"net/http"
	"sync"

	"github.com/bytedance/sonic"
	iam_domain "github.com/digiconvent/d9t/pkg/iam/domain"
	"github.com/julienschmidt/httprouter"
)

type Handler func(*Context) error

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  httprouter.Params

	Author *iam_domain.User
}

var contextPool = sync.Pool{
	New: func() any {
		return &Context{}
	},
}

func reserveContext(out http.ResponseWriter, in *http.Request, params httprouter.Params) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Writer = out
	ctx.Request = in
	ctx.Params = params

	ctx.Author = nil // always reset author

	return ctx
}

func freeContext(ctx *Context) {
	ctx.Writer = nil
	ctx.Request = nil
	ctx.Params = nil
	ctx.Author = nil
}

func (c *Context) JSON(status int, v any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	bytes, err := sonic.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.Writer.Write(bytes)
	return err
}

func (c *Context) BindJSON(v any) error {
	contents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(contents, v)
}
