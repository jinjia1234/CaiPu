package service

import (
	"CaiPu/config"
	"CaiPu/dao"
)

type Service struct {
	cfg *config.Configuration
	dao *dao.Dao
}

func New(c *config.Configuration) *Service {
	return &Service{cfg: c, dao: dao.New()}
}
