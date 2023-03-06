package view

import (
	"book/db/init"
	"fmt"
	"log"
)

type User struct {
	UserName string
	PassWord string
}

//AddUser 新增用户
func AddUser(user User) {
	tx, err := init.DB.Begin()
	if err != nil {
		log.Fatal("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`,`password`) VALUE (?,?)")
	if err != nil {
		log.Fatal("prepare fail")
		return
	}

	// 传参到sql中执行
	res, err := stmt.Exec(user.UserName, user.PassWord)
	if err != nil {
		fmt.Println("exec fail")
		return
	}
	// 提交
	tx.Commit()
	log.Println(res.LastInsertId())
	return
}
