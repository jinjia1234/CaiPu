package model

import "time"

type Food struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`             // 名称
	State     string    `json:"state"`            // 状态
	BianMa    string    `gorm:"column:BianMa"`    // 编码
	ChangJing string    `gorm:"column:ChangJing"` // 场景
	LeiXing   string    `gorm:"column:LeiXing"`   // 类型
	FangShi   string    `gorm:"column:FangShi"`   // 方式
	ShiCai    string    `gorm:"column:ShiCai"`    // 食材
	TeSe      string    `gorm:"column:TeSe"`      // 特色
	DengJi    string    `gorm:"column:DengJi"`    // 等级
	TuiJian   string    `gorm:"column:TuiJian"`   // 推荐
	BeiZhu    string    `gorm:"column:BeiZhu"`    // 备注
	CreatedAt time.Time `gorm:"created_at"`       // 创建时间
	UpdateAt  time.Time `gorm:"update_at"`        // 更新时间
}

type Foods struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`       // 名称
	State     string    `db:"state"`      // 状态
	BianMa    string    `db:"BianMa"`     // 编码
	ChangJing string    `db:"ChangJing"`  // 场景
	LeiXing   string    `db:"LeiXing"`    // 类型
	FangShi   string    `db:"FangShi"`    // 方式
	ShiCai    string    `db:"ShiCai"`     // 食材
	TeSe      string    `db:"TeSe"`       // 特色
	DengJi    string    `db:"DengJi"`     // 等级
	TuiJian   string    `db:"TuiJian"`    // 推荐
	BeiZhu    string    `db:"BeiZhu"`     // 备注
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdateAt  time.Time `db:"update_at"`  // 更新时间

}

func (Food) TableName() string {
	return "food"
}
