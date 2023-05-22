package user

import (
	"book/initalize/database/mysql"
	"log"
)

type Address struct {
	AddressID int    `json:"address_id"`
	Addressee string `json:"addressee"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

// AddUserAddress 添加用户地址
func AddUserAddress(requestsData Address, uuid string) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO address (`uuid`,`addressee`,`telephone`,`address`) VALUE (?,?,?,?)")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(uuid, requestsData.Addressee, requestsData.Telephone, requestsData.Address)
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

// GetUserAllAddress 获取用户所有地址信息
func GetUserAllAddress(uuid string) (result []Address, err error) {
	rows, err := mysql.Mysql().DB.Query("SELECT address_id ,addressee,telephone,address FROM address WHERE uuid = ? ", uuid)
	if err != nil {
		log.Println(err)
		return result, err
	}
	for rows.Next() {
		var f Address
		err = rows.Scan(&f.AddressID, &f.Addressee, &f.Telephone, &f.Address)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

// GetUserAddress 获取指定uuid对应的 addressID的地址
func GetUserAddress(uuid, AddressId string) (result Address, err error) {
	err = mysql.Mysql().DB.QueryRow("SELECT address_id,addressee,telephone,address FROM address WHERE address_id = ? and uuid=? ",
		AddressId, uuid).Scan(&result.AddressID, &result.Addressee, &result.Telephone, &result.Address)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return result, err
}

func UpdateUserAddress(requestsData Address, uuid string) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return err
	}

	// 准备sql语句
	stmt, err := tx.Prepare("UPDATE address SET addressee=?,telephone=?,address=? WHERE address_id=? and uuid=?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(requestsData.Addressee, requestsData.Telephone, requestsData.Address, requestsData.AddressID, uuid)
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

func DeleteUserAddress(AddressID, uuid string) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM address WHERE uuid = ? and address_id=?")
	if err != nil {
		log.Println("prepare fail")
		return err
	}

	// 传参到sql中执行
	_, err = stmt.Exec(uuid, AddressID)
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
