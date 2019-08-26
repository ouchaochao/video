package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video/api/handlers"
	"video/api/session"
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
	router.GET("/user/:username", handlers.GetUserInfo)
	router.POST("/user/:username/videos", handlers.AddNewVideo)
	router.GET("/user/:username/videos", handlers.ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", handlers.DeleteVideo)
	router.POST("/videos/:vid-id/comments", handlers.PostComment)
	router.GET("/videos/:vid-id/comments", handlers.ShowComments)
	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

/*
流程
main->middleware->defs(message, err)->handlers->dbops->response
*/
func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
