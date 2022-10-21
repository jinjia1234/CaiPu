package service

import (
	"CaiPu/model"
	log "github.com/sirupsen/logrus"
)

// FindFoodById 根据菜ID查询菜
func (sv *Service) FindFoodById(id int64) (bool, model.Food, error) {
	var (
		food model.Food
		exit bool
		err  error
	)
	food, err = sv.dao.FindFoodByID(id)
	if err != nil {
		log.Errorf("sv.dao.FindFoodByID(%d) error(%v)", id, err)
		return exit, food, err
	}
	if food.ID == id && id != 0 {
		exit = true
	}
	return exit, food, nil
}
