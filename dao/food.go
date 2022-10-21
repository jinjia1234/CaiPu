package dao

import (
	"CaiPu/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

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
