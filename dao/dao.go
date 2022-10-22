package dao

import (
	"CaiPu/config"
	"CaiPu/library/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New(c *config.Configuration) *Dao {
	return &Dao{
		db: sql.NewMysql(c.Database),
	}
}

func (d *Dao) Begin() *gorm.DB {
	return d.db.Begin()
}
