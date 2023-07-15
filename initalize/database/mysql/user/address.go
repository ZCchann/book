package user

import (
	"book/initalize/database/mysql"
	"fmt"
)

type Address struct {
	AddressID int    `json:"address_id"`
	Addressee string `json:"addressee"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
	UUID      string `json:"uuid"`
}

// AddUserAddress 添加用户地址
func AddUserAddress(requestsData Address, uuid string) (err error) {
	requestsData.UUID = uuid
	err = mysql.Mysql().Table("address").Create(&requestsData).Error
	if err != nil {
		err = fmt.Errorf("AddUserAddress插入数据错误 请检查: %s", err)
		return
	}
	return nil
}

// GetUserAllAddress 获取用户所有地址信息
func GetUserAllAddress(uuid string) (result []Address, err error) {
	err = mysql.Mysql().Table("address").Where("uuid=?", uuid).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetUserAllAddress 查询address表错误 请检查: %s", err)
		return
	}
	return
}

// GetUserAddress 获取指定uuid对应的 addressID的地址
func GetUserAddress(uuid, AddressId string) (result Address, err error) {
	err = mysql.Mysql().Table("address").Where("uuid=? AND address_id=?", uuid, AddressId).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetUserAddress 查询address表错误 请检查: %s", err)
		return
	}
	return

}

func UpdateUserAddress(requestsData Address, uuid string) (err error) {
	err = mysql.Mysql().Table("address").Where("uuid=? AND address_id=?", uuid, requestsData.AddressID).Save(&requestsData).Error
	if err != nil {
		err = fmt.Errorf("UpdateUserAddress update错误 请检查: %s", err)
		return
	}
	return
}

func DeleteUserAddress(AddressID, uuid string) (err error) {
	err = mysql.Mysql().Table("address").Where("uuid=? AND address_id=?", uuid, AddressID).Delete(&Address{}).Error
	if err != nil {
		err = fmt.Errorf("address删除数据错误 请检查: %s", err)
		return
	}
	return nil
}
