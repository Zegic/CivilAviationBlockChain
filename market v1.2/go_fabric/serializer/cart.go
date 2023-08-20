package serializer

import (
	"context"
	"go_fabric/conf"
	"go_fabric/dao"
	"go_fabric/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductName   string `json:"product_name"`
	ProductId     uint   `json:"product_id"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.Boss) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		ProductName:   product.Name,
		BossID:        boss.ID,
		BossName:      boss.BossName,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: product.DiscountPrice,
		OnSale:        product.OnSale,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewBossDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetBossById(item.BossId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
