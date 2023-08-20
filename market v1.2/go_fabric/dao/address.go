package dao

import (
	"context"
	"go_fabric/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(address model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

func (dao *AddressDao) ShowAddress(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}

func (dao *AddressDao) ListAddress(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&addresses).Error
	return
}

func (dao *AddressDao) UpdateAddress(aId uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id=?", aId).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddress(aId uint, uId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id=? AND user_id =?", aId, uId).Delete(&model.Address{}).Error
}
