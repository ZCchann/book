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
	Permission bool   `json:"Permission"`
	User       bool   `json:"user"`
	RuleName   string `json:"rulename"`
}
