package server

import "github.com/julienschmidt/httprouter"

type Router struct {
	Router *httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		Router: httprouter.New(),
	}
}
