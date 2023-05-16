package authority

import "database/sql"

type Permission struct {
	Name  string `json:"name"`
	State bool   `json:"state"`
}
type EditPermissions struct {
	ID          int          `json:"id"`
	RuleName    string       `json:"rule_name"`
	Permissions []Permission `json:"permissions"`
}

type Column struct {
	Name       string
	Type       string
	Null       string
	Key        string
	Default    sql.NullString
	Extra      string
	Privileges string
}

type Authority struct {
	ID                   int    `json:"id"`
	DataManagement       bool   `json:"data_management"`
	OrderManagement      bool   `json:"order_management"`
	PermissionManagement bool   `json:"permission_management"`
	UserManagement       bool   `json:"user_management"`
	RuleName             string `json:"rulename"`
}
