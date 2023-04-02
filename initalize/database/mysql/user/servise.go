package user

import (
	"book/initalize/database/mysql"
	"fmt"
	"log"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UUID     string `json:"uuid"`
}

//GetUser 通过用户名获取用户信息
func GetUser(UserName string) (u User, err error) {
	var data User
	err = mysql.Mysql().DB.QueryRow("SELECT username,password,uuid FROM user WHERE username = ? ", UserName).Scan(&data.Username, &data.Password, &data.UUID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return data, err

}

//GetUserForID 通过ID获取用户信息 不含密码
func GetUserForID(uuid string) (data User, err error) {
	err = mysql.Mysql().DB.QueryRow("SELECT uuid,username,email FROM user WHERE uuid = ? ", uuid).Scan(&data.UUID, &data.Username, &data.Email)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return data, err

}

//GetAllUser 获取所有用户的用户名
func GetAllUser() (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT id,username,email,uuid FROM user")
	if err != nil {
		log.Println(err)
		return result, err
	}
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Id, &f.Username, &f.Email, &f.UUID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

//AddUser 新增用户
func AddUser(UserName, Password, Email string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`,`password`,`email`,`uuid`) VALUE (?,?,?,replace(uuid(),\"-\",\"\"))")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName, Password, Email)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return err
	}
	return nil
}

func DelUser(UUID string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE uuid = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UUID)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return err
	}
	return nil
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(uuid, Email, Password string) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET password = ?, email= ? WHERE uuid = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(Password, Email, uuid)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return err
	}
	return nil
}

func SearchUser(username string) (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query(fmt.Sprintf("SELECT id,username,email from user where username REGEXP '%s';", username))
	if err != nil {
		log.Println(err)
		return result, err
	}
	_ = rows.Scan()
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Id, &f.Username, &f.Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

func UpdateUserEmail(uuid, email string) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return err
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET email = ? WHERE uuid = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(email, uuid)
	if err != nil {
		log.Println("exec fail")
		return err
	}
	// 提交
	err = tx.Commit()
	if err != nil {
		log.Println("commit error ", err)
		return err
	}
	return nil
}
