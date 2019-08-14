package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video/api/defs"
	"video/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	// 千万不要用+号来连接query的各个部分, 不安全, 容易被撞库攻击
	//Prepare预编译, 更安全了, 会拦下撞库攻击
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?,  ?)")
	if err != nil {
		return err
	}
	// 执行, 将两个参数传到上面两个问号处
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	// defer 是栈退出的时候才调用, 性能会有些许损耗
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		// string 默认是没有内容的
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//Create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	// 时间格式, go的彩蛋
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, Author: aid, Name: name, DisplayCting: ctime}
	defer stmtIns.Close()
	return res, nil
}
