package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video/api/dbops"
	"video/api/defs"
	"video/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//创建用户的时候,r.Body:读取post里的表单
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	//Unmarshal:反序列化,把json转化为struct
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	//Marshal:序列化
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
