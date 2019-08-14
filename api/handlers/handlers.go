package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//dbops.AddUserCredential("aaa", "bbb")
	//pwd, _ := dbops.GetUserCredential("aaa")
	//dbops.DeleteUser("aaa","bbb")
	//res, _ :=dbops.AddNewVideo(123, "kk1")
	//dbops.GetVideoInfo("123")
	//dbops.DeleteVideoInfo("123")
	//dbops.AddNewComments("1",123,"test")
	//dbops.ListComments("1",)
	io.WriteString(w, "ok")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	io.WriteString(w, uname)
}
