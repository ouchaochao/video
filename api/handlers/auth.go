/*
文件初始化
 */
package handlers

import (
	"net/http"
	"video/api/defs"
	"video/api/session"
)

//X开头的都是自定义的header,加到原生http header调用里面,构成整个健全过程
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FAILD_UNAME = "X-User-Name"

//检测用户session是否合法，是就返回username
func ValidateUserSession(r *http.Request) bool {
	//get获取sessionId
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	//检测过期否?
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FAILD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FAILD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
