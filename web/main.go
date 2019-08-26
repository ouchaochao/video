package main

import (
	"net/http"
	// "html/template"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.HomeHandler)
	router.POST("/", handlers.HomeHandler)
	router.GET("/userhome", handlers.UserHomeHandler)
	router.POST("/userhome", handlers.UserHomeHandler)
	router.POST("/api", handlers.ApiHandler)
	router.POST("/upload/:vid-id", handlers.ProxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
