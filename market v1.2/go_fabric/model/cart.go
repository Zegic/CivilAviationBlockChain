package model

import "gorm.io/gorm"

// 购物车
type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       uint `gorm:"not null"` //数量
	MaxNum    uint `gorm:"not null"` //限购数量
	Check     bool //是否支付
}
