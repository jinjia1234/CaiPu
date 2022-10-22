package service

import (
	"CaiPu/config"
	"CaiPu/dao"
	"CaiPu/model"
)

//type Service struct {
//	cfg *config.Configuration
//	dao *dao.Dao
//}

func New(c *config.Configuration) {

	dao.New(c)

}

func QueryMultiRowDemo() []model.Foods {
	demo := dao.QueryMultiRowDemo()

	return demo
}
