package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video/api/defs"
	"video/api/utils"
)

/*
下面三个函数分别是：
	添加用户
	获取用户信息
	删除用户
*/
func AddUserCredential(loginName string, pwd string) error {
	// 千万不要用+号来连接query的各个部分, 不安全, 容易被撞库攻击
	//Prepare预编译, 更安全了, 会拦下撞库攻击
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
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

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var id int
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}
	defer stmtOut.Close()
	return res, nil
}

/*
下面三个函数分别是：
	添加视频
	获取视频
	删除视频
*/
func AddVideoInfo(aid int, name string) (*defs.VideoInfo, error) {
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
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	//怎么测试：res.XX
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	//var不能写在一起
	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return nil
	}
	defer stmtDel.Close()
	return nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name=? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time<=FROM_UNIXTIME(?)
		OREDER BY video_info.create_time DESC`)
	var res []*defs.VideoInfo
	if err != nil {
		return res, err
	}
	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}
		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}
	defer stmtOut.Close()
	return res, nil
}

/*
下面两个函数分别是：
	添加评论
	查看评论
	此处不添加删除评论
*/
func AddComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return nil
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	// 连接user和comments表查询字段
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id=users.id
		WHERE comments.video_id=? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	//此处定义了Comment，放在了apidef.go文件中
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
