package user

import (
	"book/initalize/database/mysql"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func GetUser(UserName string) (u User, err error) {
	var data User
	err = mysql.Mysql().DB.QueryRow("SELECT username,password FROM user WHERE username = ? ", UserName).Scan(&data.Username, &data.Password)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return data, err

}

//GetAllUser 获取所有用户的用户名
func GetAllUser() (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT username,email FROM user")
	if err != nil {
		log.Println(err)
		return result, err
	}
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Username, &f.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

//AddUser 新增用户
func AddUser(UserName, Password string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`,`password`) VALUE (?,?)")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName, Password)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
	}
	return nil
}

func DelUser(UserName string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
	}
	return nil
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(UserName, Password string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET password = ? WHERE username = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(Password, UserName)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
	}
	return nil
}
