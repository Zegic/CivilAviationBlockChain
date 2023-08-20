package model

import "gorm.io/gorm"

// 收藏夹
type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      Boss    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}
