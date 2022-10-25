package model

import "time"

type FoodMetaSave struct {
	Field1 string `json:"a0"`
	Field2 string `json:"a1"`
	Field3 string `json:"a2"`
	Field4 string `json:"a3"`
	Field5 string `json:"a4"`
	Field6 string `json:"a5"`
	Field7 string `json:"a6"`
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
