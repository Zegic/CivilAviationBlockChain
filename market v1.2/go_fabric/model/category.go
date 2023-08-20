package model

import "gorm.io/gorm"

// 种类
type Category struct {
	gorm.Model
	CategoryName string
}
