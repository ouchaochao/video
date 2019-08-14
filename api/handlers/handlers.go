package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"video/api/dbops"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
	dbops.AddUserCredential("aaa", "bbb")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	io.WriteString(w, uname)
}
