package model

import "time"

type FoodMetaSave struct {
	ChangJing []string `json:"changjing"`
	LeiXing   []string `json:"leixing"`
	FangShi   []string `json:"fangshi"`
	TeSe      []string `json:"tese"`
	DengJi    []string `json:"dengji"`
	ZhiShu    []string `json:"zhishu"`
	Ting      []string `json:"ting"`
}

type FoodMetaTest struct {
	Data []string `json:"data"`
}

type FoodMeta struct {
	ID        int64     `db:"id"`
	MetaKey   string    `db:"meta_key"`
	MetaValue string    `db:"meta_value"`
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdateAt  time.Time `db:"update_at"`  // 更新时间
}

func (FoodMeta) TableName() string {
	return "food_meta"
}
