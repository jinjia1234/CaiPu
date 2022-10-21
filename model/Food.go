package model

import "time"

type Food struct {
	ID         int64     `json:"id"`
	Type       string    `json:"type"` // 类型
	Name       string    `json:"name"` // 名称
	CreateTime time.Time `json:"created_at"`
	UpdateTime time.Time `json:"update_at"`
}

func (Food) TableName() string {
	return "food"
}
