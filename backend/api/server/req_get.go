package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (r *Router) Get(path string, handler Handler) {
	r.Router.GET(path, func(out http.ResponseWriter, in *http.Request, p httprouter.Params) {
		context := reserveContext(out, in, p)

		defer freeContext(context)
		err := handler(context)
		if err != nil {
			r.handleError(context, err)
		}
	})
}

func (r *Router) handleError(context *Context, err error) {
	fmt.Println(context)
	fmt.Println(err)
}

func (r *Router) ServeHttps() {
}
