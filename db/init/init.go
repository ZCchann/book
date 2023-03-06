package init

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	IP       = "192.168.1.201"
	PORT     = "3306"
	dbName   = "book"
)

var DB *sql.DB

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")

}
