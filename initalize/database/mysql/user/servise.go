package user

import (
	"book/initalize/database/mysql"
	"fmt"
)

//GetUser 通过用户名获取用户信息
func GetUser(UserName string) (data User, err error) {
	//var data User
	//err = mysql.Mysql().DB.QueryRow("SELECT username,password,uuid FROM user WHERE username = ? ", UserName).Scan(&data.Username, &data.Password, &data.UUID)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//return data, err

	err = mysql.Mysql().Select("username,password,uuid").Table("user").Where("username=?", UserName).Find(&data).Error
	if err != nil {
		err = fmt.Errorf("GetUser 读取user表错误 请检查: %s", err)
		return User{}, err
	}
	return

}

//GetUserForID 通过ID获取用户信息 不含密码
func GetUserForID(uuid string) (data User, err error) {
	//err = mysql.Mysql().DB.QueryRow("SELECT uuid,username,email,authorityID FROM user WHERE uuid = ? ", uuid).Scan(&data.UUID, &data.Username, &data.Email, &data.AuthorityID)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//return data, err
	err = mysql.Mysql().Table("user").Select("uuid,username,email,authorityID").Where("uuid=?", uuid).Find(&data).Error
	if err != nil {
		err = fmt.Errorf("GetUserForID 读取user表错误 请检查: %s", err)
		return User{}, err
	}
	return
}

//GetAllUser 获取所有用户的用户名
func GetAllUser() (result []User, err error) {
	err = mysql.Mysql().Table("user").Select("id,username,email,uuid,authorityID").Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetAllUser 读取user表错误 请检查: %s", err)
		return
	}
	return
}

//AddUser 新增用户
func AddUser(request User) (err error) {
	err = mysql.Mysql().Table("user").Create(&request).Error
	if err != nil {
		err = fmt.Errorf("bookdata插入数据错误 请检查: %s", err)
		return
	}
	return nil
}

// DelUser 删除用户
func DelUser(UUID string) (err error) {
	err = mysql.Mysql().Table("user").Where("uuid=?", UUID).Delete(&User{}).Error
	if err != nil {
		err = fmt.Errorf("user表删除数据错误 请检查: %s", err)
		return
	}
	return nil
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(uuid, Email, Password, Form string, authorityID int) (err error) {
	if Form == "admin" {
		err = mysql.Mysql().Table("user").Where("uuid = ?", uuid).Updates(map[string]interface{}{
			"password":    Password,
			"authorityID": authorityID,
			"email":       Email,
		}).Error
		if err != nil {
			err = fmt.Errorf("UpdateUserPassword 更新user表错误 请检查: %s", err)
			return
		}
	} else {
		//	stmt, err := tx.Prepare("UPDATE user SET email = ?  WHERE uuid = ?")
		err = mysql.Mysql().Table("user").Where("uuid = ?", uuid).Updates(map[string]interface{}{
			"password": Password,
			"email":    Email,
		}).Error
		if err != nil {
			err = fmt.Errorf("UpdateUserPassword 更新user表错误 请检查: %s", err)
			return
		}
	}
	return
}

// SearchUser 通过用户名搜索用户信息
func SearchUser(username string) (result []User, err error) {
	err = mysql.Mysql().Table("user").Select("id,username,email,uuid ,authorityID").Where("username REGEXP ?", username).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("SearchUser 读取user表错误 请检查: %s", err)
		return
	}
	return
}

// UpdateUserEmail 更新用户信息
func UpdateUserEmail(uuid, email, Form string, authorityID int) (err error) {
	if Form == "admin" {
		err = mysql.Mysql().Table("user").Where("uuid = ?", uuid).Updates(map[string]interface{}{
			"email":       email,
			"authorityID": authorityID,
		}).Error
		if err != nil {
			err = fmt.Errorf("UpdateUserEmail 更新user表错误 请检查: %s", err)
			return
		}
	} else {
		//	stmt, err := tx.Prepare("UPDATE user SET email = ?  WHERE uuid = ?")
		err = mysql.Mysql().Table("user").Where("uuid = ?", uuid).Updates(map[string]interface{}{
			"email": email,
		}).Error
		if err != nil {
			err = fmt.Errorf("UpdateUserEmail 更新user表错误 请检查: %s", err)
			return
		}
	}
	return
}

// GetUserPermissionForID 通过权限ID查询用户
func GetUserPermissionForID(PermissionID int) (result []User, err error) {
	err = mysql.Mysql().Table("user").Select("id,username,email,uuid,authorityID").Where("authorityID=?", PermissionID).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetUserPermissionForID 读取user表错误 请检查: %s", err)
		return
	}
	return

}
