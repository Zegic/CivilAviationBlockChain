package dao

import (
	"context"
	"go_fabric/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&cart).Error
}

func (dao *CartDao) UpdateCartNum(uId uint, cId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).Where("id=? AND user_id = ?", cId, uId).Update("num", num).Error
}

func (dao *CartDao) DeleteCart(cId uint, uId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=? AND user_id =?", cId, uId).Delete(&model.Cart{}).Error
}

func (dao *CartDao) ListCart(uId uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&carts).Error
	return
}
