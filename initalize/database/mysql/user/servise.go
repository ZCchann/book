package user

import (
	"book/initalize/database/mysql"
	"fmt"
	"log"
)

type User struct {
	UserName string
	PassWord string
}

func GetUser(UserName string) (username string, err error) {
	var name string
	err = mysql.Mysql().DB.QueryRow("SELECT username FROM user WHERE username = ? ", UserName).Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	return name, err

}

//AddUser 新增用户
func AddUser(UserName, Password string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Fatal("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`,`password`) VALUE (?,?)")
	if err != nil {
		log.Fatal("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName, Password)
	if err != nil {
		fmt.Println("exec fail")
		return err
	}
	// 提交
	tx.Commit()
	return nil
}

func DelUser(UserName string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Fatal("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		log.Fatal("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName)
	if err != nil {
		fmt.Println("exec fail")
		return err
	}
	// 提交
	tx.Commit()
	return nil
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(UserName, Password string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Fatal("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET password = ? WHERE username = ?")
	if err != nil {
		log.Fatal("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(Password, UserName)
	if err != nil {
		fmt.Println("exec fail")
		return err
	}
	// 提交
	tx.Commit()
	return nil
}
