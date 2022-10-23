package service

import (
	"CaiPu/config"
	"CaiPu/dao"
	"CaiPu/library/sql"
	"github.com/jinzhu/gorm"
)

type Service struct {
	cfg *config.Configuration
	dao *dao.Dao
	db  *gorm.DB
}

func New(c *config.Configuration) *Service {
	return &Service{
		cfg: c,
		dao: dao.New(c),
		db:  sql.NewMysql(c.Database),
	}
}
