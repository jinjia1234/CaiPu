package dbsql

import (
	"CaiPu/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DbInit = DbWork{}

type DbWork struct {
	Conn *sqlx.DB
}

func New(c *config.Configuration) {
	s := fmt.Sprintf("%s:%s@(%s:%s)/%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.SchemaName)
	db, err := sqlx.Connect("mysql", s+"?parseTime=true&loc=Asia%2FShanghai&charset=utf8")
	if err != nil {
		fmt.Println("open mysql err %s", err)
	}
	db.SetMaxOpenConns(20)
	DbInit.Conn = db
}
