package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video/api/handlers"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	handlers.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:username", handlers.Login)
	return router
}

/*
流程
main->middleware->defs(message, err)->handlers->dbops->response
*/
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
