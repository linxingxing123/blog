package models

import "gin_blog/dao"

// 定义 模型 与 数据库中的表相对应

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	CreateTime int64
}

//--------------数据库操作-----------------

// InsertUser 插入
func InsertUser(user *User) (int64, error) {
	return dao.ModifyDB("insert into users(username,password,status,create_time) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.CreateTime)
}

// QueryUserWithUsername 根据用户名查询id
func QueryUserWithUsername(username string) int {
	var user User
	err := dao.QueryRowDB(&user, "select id from users where username=?", username)
	if err != nil {
		return 0
	}
	return user.Id
}

// QueryUserWithParam 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	var user User
	err := dao.QueryRowDB(&user, "select id from users where username=? and password=?", username, password)
	if err != nil {
		return 0
	}
	return user.Id
}
