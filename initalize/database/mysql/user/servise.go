package user

import (
	"book/initalize/database/mysql"
	"fmt"
	"log"
)

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	UUID        string `json:"uuid"`
	AuthorityID int    `json:"authorityid"`
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
	err = mysql.Mysql().DB.QueryRow("SELECT uuid,username,email,authorityID FROM user WHERE uuid = ? ", uuid).Scan(&data.UUID, &data.Username, &data.Email, &data.AuthorityID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return data, err

}

//GetAllUser 获取所有用户的用户名
func GetAllUser() (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT id,username,email,uuid ,authorityID FROM user")
	if err != nil {
		log.Println(err)
		return result, err
	}
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Id, &f.Username, &f.Email, &f.UUID, &f.AuthorityID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

//AddUser 新增用户
func AddUser(UserName, Password, Email string, authorityID int) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`,`password`,`email`,`uuid`,`authorityID`) VALUE (?,?,?,replace(uuid(),\"-\",\"\"),?)")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(UserName, Password, Email, authorityID)
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

// DelUser 删除用户
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
func UpdateUserPassword(uuid, Email, Password string, authorityID int) error {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET password = ?, email= ? ,authorityID= ? WHERE uuid = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(Password, Email, authorityID, uuid)
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

// SearchUser 通过用户名搜索用户信息
func SearchUser(username string) (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query(fmt.Sprintf("SELECT id,username,email,uuid ,authorityID from user where username REGEXP '%s';", username))
	if err != nil {
		log.Println(err)
		return result, err
	}
	_ = rows.Scan()
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Id, &f.Username, &f.Email, &f.UUID, &f.AuthorityID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

// UpdateUserEmail 更新用户信息
func UpdateUserEmail(uuid, email string, authorityID int) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return err
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET email = ? ,authorityID = ? WHERE uuid = ?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(email, authorityID, uuid)
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

// GetUserPermissionForID 通过权限ID查询用户
func GetUserPermissionForID(PermissionID int) (result []User, err error) {
	rows, err := mysql.Mysql().DB.Query(fmt.Sprintf("SELECT id, username, email, uuid, authorityID FROM user WHERE authorityID = %d", PermissionID))
	if err != nil {
		log.Println(err)
		return result, err
	}
	for rows.Next() {
		var f User
		err = rows.Scan(&f.Id, &f.Username, &f.Email, &f.UUID, &f.AuthorityID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}
