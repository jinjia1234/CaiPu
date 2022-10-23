package service

import (
	"CaiPu/model"
	log "github.com/sirupsen/logrus"
)

// CreateFood 创建菜
func (sv *Service) CreateFood(d model.Food) (int64, error) {
	id, err := sv.dao.CreateFood(d)
	if err != nil {
		return id, err
	}
	return id, nil
}

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

// FindFoods 查询菜列表
func (sv *Service) FindFoods() ([]model.Food, error) {
	var (
		err   error
		foods []model.Food
	)
	foods, err = sv.dao.FindFoods()
	if err != nil {
		println(err.Error())
		log.Errorf("sv.dao.FindFoods() error(%v)", err)
		return foods, err
	}
	return foods, nil
}

// CountFood 统计
func (sv *Service) CountFood() (int64, int64, int64, error) {
	var (
		err                                error
		countFood, countHunCai, countSuCai int64
	)
	if err = sv.db.Model(model.Food{}).Count(&countFood).Error; err != nil {
		log.Errorf("db.Model(Food).Count() error(%v)", err)
		return countFood, countHunCai, countSuCai, nil
	}
	if err = sv.db.Model(model.Food{}).
		Where("LeiXing = ?", "荤菜").
		Count(&countHunCai).Error; err != nil {
		log.Errorf("db.Model(Food).Count() error(%v)", err)
		return countFood, countHunCai, countSuCai, nil
	}
	if err = sv.db.Model(model.Food{}).
		Where("LeiXing = ?", "素菜").
		Count(&countSuCai).Error; err != nil {
		log.Errorf("db.Model(Food).Count() error(%v)", err)
		return countFood, countHunCai, countSuCai, nil
	}
	return countFood, countHunCai, countSuCai, nil
}
