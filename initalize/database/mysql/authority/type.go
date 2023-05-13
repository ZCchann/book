package authority

import "database/sql"

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
	Data       int    `json:"data"`
	Order      int    `json:"order"`
	Permission int    `json:"permission"`
	User       int    `json:"user"`
	RuleName   string `json:"rulename"`
}
