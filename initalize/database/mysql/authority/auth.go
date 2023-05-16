package authority

import (
	"book/initalize/database/mysql"
	"log"
)

// GetRuleForUUID 通过UUID 查询权限详情
func GetRuleForUUID(uuid string) (result Authority, err error) {
	err = mysql.Mysql().DB.QueryRow("SELECT authority.data_management, authority.order_management, authority.permission_management, authority.user_management FROM user JOIN authority ON user.authorityID = authority.id WHERE user.uuid = ?;", uuid).Scan(&result.DataManagement, &result.OrderManagement, &result.PermissionManagement, &result.UserManagement)
	if err != nil {
		return
	}
	return result, err
}

// GetPermissionsGroup 所有权限细节信息
func GetPermissionsGroup() (columns []Column, err error) {
	rows, err := mysql.Mysql().DB.Query("SHOW COLUMNS FROM authority;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var column Column
		if err := rows.Scan(&column.Name, &column.Type, &column.Null, &column.Key, &column.Default, &column.Extra); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

// GetAllPermissionsIDName 返回所有的权限组ID、权限名称
func GetAllPermissionsIDName() (result []EditPermissions, err error) {
	rows, err := mysql.Mysql().DB.Query("select id,rulename from authority;")
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var f EditPermissions
		err = rows.Scan(&f.ID, &f.RuleName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, err
}

// GetPermissionsByID 通过ID获取权限详情
func GetPermissionsByID(ID string) (result Authority, err error) {
	err = mysql.Mysql().DB.QueryRow("select * from authority where id= ?", ID).Scan(&result.ID, &result.DataManagement, &result.OrderManagement, &result.PermissionManagement, &result.UserManagement, &result.RuleName)
	if err != nil {
		log.Println(err)
		return
	}
	return result, err
}

// UpdatePermissionsByID 根据权限ID更改内容
func UpdatePermissionsByID(data EditPermissions) (err error) {
	tx, err := mysql.Mysql().DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return err
	}
	var DataManagement, OrderManagement, permissionManagement, userManagement bool

	for _, i := range data.Permissions {
		log.Println("name ", i.Name, " status ", i.State)
		switch i.Name {
		case "DataManagement":
			DataManagement = i.State
		case "OrderManagement":
			OrderManagement = i.State
		case "PermissionManagement":
			permissionManagement = i.State
		case "UserManagement":
			userManagement = i.State
		}
	}

	stmt, err := tx.Prepare("UPDATE authority SET data_management=?,order_management=?,permission_management=?,user_management=? WHERE id=?")
	if err != nil {
		log.Println("Prepare fail ", err)
		return err
	}
	_, err = stmt.Exec(DataManagement, OrderManagement, permissionManagement, userManagement, data.ID)
	if err != nil {
		log.Println("exec fail ", err)
		return err
	}
	tx.Commit()
	return nil
}
