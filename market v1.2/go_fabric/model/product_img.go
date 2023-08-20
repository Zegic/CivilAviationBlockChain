package model

import "gorm.io/gorm"

// 商品图片
type ProductImg struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}
