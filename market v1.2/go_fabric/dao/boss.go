package dao

import (
	"context"
	"go_fabric/model"
	"gorm.io/gorm"
)

type BossDao struct {
	*gorm.DB
}

func NewBossDao(ctx context.Context) *BossDao {
	return &BossDao{NewDBClient(ctx)}
}

func NewBossDaoDB(db *gorm.DB) *BossDao {
	return &BossDao{db}
}

// 根据商户名是否存在该商户
func (dao *BossDao) ExistOrNotByBossName(bossName string) (boss *model.Boss, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Boss{}).Where("boss_name=?", bossName).Find(&boss).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return boss, true, nil
}

// 根据ID寻找商户
func (dao *BossDao) GetBossById(id uint) (boss *model.Boss, err error) {
	err = dao.DB.Model(&model.Boss{}).Where("id=?", id).First(&boss).Error
	return
}

// 商户列表
func (dao *BossDao) ListBoss(pageNum int, pageSize int) (bosses []model.Boss, err error) {
	offset := (pageNum - 1) * pageSize
	err = dao.DB.Model(&model.Boss{}).Offset(offset).Limit(pageSize).Find(&bosses).Error
	return
}

// 删除商户
func (dao *BossDao) DeleteByBossName(boss *model.Boss) error {
	return dao.DB.Delete(&boss).Error
}

// 封禁商户
func (dao *BossDao) PassiveByUserName(uId uint) error {
	return dao.DB.Model(&model.Boss{}).Where("id=?", uId).Update("status", "passive").Error
}
