package authority

import (
	"book/initalize/database/mysql"
	"log"
)

func GetRule(UserName string) (result Auth, err error) {
	err = mysql.Mysql().DB.QueryRow("SELECT admin, order, rulename FROM user JOIN permissions ON user.authorityID = id WHERE user.username = ?;", UserName).Scan(&result.Admin, &result.Order, &result.RuleName)
	if err != nil {
		log.Println(err)
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
	err = mysql.Mysql().DB.QueryRow("select * from authority where id= ?", ID).Scan(&result.ID, &result.Data, &result.Order, &result.Permission, &result.User, &result.RuleName)
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
		case "Data":
			DataManagement = i.State
		case "Order":
			OrderManagement = i.State
		case "Permission":
			permissionManagement = i.State
		case "User":
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
