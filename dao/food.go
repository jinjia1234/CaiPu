package dao

import (
	"CaiPu/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// CreateFood 创建菜
func (d *Dao) CreateFood(food model.Food) (int64, error) {
	var (
		err error
	)
	if err = d.db.Model(model.Food{}).Create(&food).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Errorf("db.Model(Food).Create(%v) error(%v)", food, err)
		return food.ID, err
	}
	return food.ID, nil
}

// FindFoodByID 根据用户id查询用户
func (d *Dao) FindFoodByID(Id int64) (model.Food, error) {
	var (
		food model.Food
		err  error
	)

	if err = d.db.Model(model.Food{}).
		Where("id = ?", Id).
		First(&food).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Errorf("db.Model(Food).First(%d) error(%v)", Id, err)
		return food, err
	}
	return food, nil
}

// FindFoods 查询所有菜
func (d *Dao) FindFoods() ([]model.Food, error) {
	var (
		foods []model.Food
		err   error
	)
	if err = d.db.Model(model.Food{}).Find(&foods).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Errorf("db.Model(Food).Find() error(%v)", err)
		return foods, err
	}
	return foods, nil
}
