package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type mdb struct {
	*sql.DB
}

var m = new(mdb)

func Mysql() *mdb {
	return m
}

func (m *mdb) InitDB(username, password, host, port, database string) {
	path := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8&parseTime=true"}, "")
	m.DB, _ = sql.Open("mysql", path)
	m.SetConnMaxLifetime(100)
	m.SetMaxIdleConns(10)
	if err := m.Ping(); err != nil {
		log.Fatal("open database fail err: ", err)
		return
	}
}
