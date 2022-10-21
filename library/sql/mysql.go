package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	Username   string `toml:"username"`
	Password   string `toml:"password"`
	Host       string `toml:"host"`
	Port       string `toml:"port"`
	SchemaName string `toml:"schema_name"`
	Args       string `toml:"args"`
}

func NewMysql(c Database) *gorm.DB {
	dbSource := c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.SchemaName + "?" + c.Args
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("Connet database error(%v)", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(false)
	return db
}
