package dao

import (
	"CaiPu/library/sql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New() *Dao {
	var database sql.Database
	database.Host = "43.136.132.176"
	database.Port = "14357"
	database.SchemaName = "37617718_caidan"
	database.Username = "37617718"
	database.Password = "37617718."
	database.Args = "charset=utf8mb4&parseTime=True&loc=Local"
	return &Dao{
		db: sql.NewMysql(database),
	}
}

func (d *Dao) Begin() *gorm.DB {
	return d.db.Begin()
}
