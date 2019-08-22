package session

import (
	"sync"
	"time"
	"video/api/dbops"
	"video/api/defs"
	"video/api/utils"
)

//sync.Map自己实现了一套线程安全的机制,优化了并发读,写的话要加锁
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

//从数据库读取sessionId到cache
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		//不对外返回东西
		return
	}
	//返回值
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}


func nowInMilli() int64 {
	//UnixNano纳秒级别, '/ 100000'后变成毫秒级别
	return time.Now().UnixNano() / 100000
}
//产生sessionId
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	//过期时间30min(单位:毫秒)
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id

}

//session过期否?
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpieredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
//删除过期的session
func deleteExpieredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
