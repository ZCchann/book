package authority

import (
	"book/initalize/database/mysql"
	"fmt"
)

// GetRuleForUUID 通过UUID 查询权限详情
func GetRuleForUUID(uuid string) (result Authority, err error) {
	err = mysql.Mysql().Joins("JOIN authority ON user.authorityID = authority.id").
		Select("authority.data_management, authority.order_management, authority.permission_management, authority.user_management").Table("user").Where("user.uuid = ?", uuid).First(&result).Error
	if err != nil {
		err = fmt.Errorf("GetRuleForUUID 查询错误 请检查: %s", err)
		return
	}
	return

}

// GetPermissionsFirst 所有数据库第一条结果
func GetPermissionsFirst() (result Authority, err error) {
	err = mysql.Mysql().Table("authority").First(&result).Error
	if err != nil {
		err = fmt.Errorf("GetPermissionsFirst 查询错误 请检查: %s", err)
		return
	}
	return
}

// GetAllPermissionsIDName 返回所有的权限组ID、权限名称
func GetAllPermissionsIDName() (result []RetPermissions, err error) {
	err = mysql.Mysql().Table("authority").Select("id,rulename").Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetAllPermissionsIDName 查询错误 请检查: %s", err)
		return
	}
	return
}

// GetPermissionsByID 通过ID获取权限详情
func GetPermissionsByID(ID string) (result Authority, err error) {
	err = mysql.Mysql().Table("authority").Where("id=?", ID).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("GetPermissionsByID 查询错误 请检查: %s", err)
		return
	}
	return

}

// AddPermission 添加权限组信息
func AddPermission(data EditPermissions) (err error) {
	var auth Authority

	for _, i := range data.Permissions {
		switch i.Name {
		case "DataManagement":
			auth.DataManagement = i.State
		case "OrderManagement":
			auth.OrderManagement = i.State
		case "PermissionManagement":
			auth.PermissionManagement = i.State
		case "UserManagement":
			auth.UserManagement = i.State
		}
	}

	err = mysql.Mysql().Table("authority").Create(&auth).Error
	if err != nil {
		err = fmt.Errorf("authority插入数据错误 请检查: %s", err)
		return
	}
	return nil

}

// UpdatePermissionsByID 根据权限ID更改内容
func UpdatePermissionsByID(data EditPermissions) (err error) {
	var auth Authority

	for _, i := range data.Permissions {
		switch i.Name {
		case "DataManagement":
			auth.DataManagement = i.State
		case "OrderManagement":
			auth.OrderManagement = i.State
		case "PermissionManagement":
			auth.PermissionManagement = i.State
		case "UserManagement":
			auth.UserManagement = i.State
		}
	}

	err = mysql.Mysql().Table("authority").Where("id=?", data.ID).Save(&auth).Error
	if err != nil {
		err = fmt.Errorf("authority 更新数据错误 请检查: %s", err)
		return
	}
	return nil

}

// DeletePermissionByID 通过权限组ID删除权限组
func DeletePermissionByID(PermissionID int) (err error) {
	var data Authority
	data.ID = PermissionID

	err = mysql.Mysql().Table("authority").Where("id=?", data.ID).Delete(&Authority{}).Error
	if err != nil {
		err = fmt.Errorf("bookdata删除数据错误 请检查: %s", err)
		return
	}
	return nil
}
