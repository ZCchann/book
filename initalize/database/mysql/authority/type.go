package authority

import (
	"book/api/permissions"
	"database/sql"
)

type Auth struct {
	Admin    int    `json:"admin"`
	Order    int    `json:"order"`
	RuleName string `json:"rulename"`
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
	ID         int    `json:"id"`
	Data       bool   `json:"data"`
	Order      bool   `json:"order"`
	Permission bool   `json:"permission"`
	User       bool   `json:"user"`
	RuleName   string `json:"rulename"`
}

type UpdatePermission struct {
	ID          int                      `json:"id"`
	Permissions []permissions.Permission `json:"permissions"`
	RuleName    string                   `json:"rule_name"`
}
