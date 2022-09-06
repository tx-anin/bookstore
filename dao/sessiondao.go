package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

//AddSession 向数据库中添加
func AddSession(s *model.Session) error {

	//sql语句
	sqlStr := "insert into sessions values (?,?,?)"

	_, err := utils.Db.Exec(sqlStr, s.SessionID, s.UserName, s.UserID)

	if err != nil {
		return err
	}

	return nil
}

//DeleteSession 删除
func DeleteSession(sID string) error {
	//sql
	sqlStr := "delete from sessions where session_id = ?;"

	_, err := utils.Db.Exec(sqlStr, sID)
	if err != nil {
		return err
	}

	return nil
}

//GetSession 根据session的ID 查询Session
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建session
	sess := &model.Session{}
	//扫描
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

//IsLogin 判断用户是否已经登录 false:没有登录 true:已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}
