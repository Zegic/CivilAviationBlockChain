package model

import (
	"go_fabric/cache"
	"gorm.io/gorm"
	"strconv"
)

// 商品
type Product struct {
	gorm.Model
	Name string `gorm:"not null"`
	//ProductId     string `gorm:""`
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string //打折后价格
	OnSale        bool   `gorm:"default:false"` //是否在售
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	//增加点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
