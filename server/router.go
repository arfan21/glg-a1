package server

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()

	return r
}
