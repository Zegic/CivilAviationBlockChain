package dao

import (
	"context"
	"go_fabric/model"
	"gorm.io/gorm"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao(ctx context.Context) *AdminDao {
	return &AdminDao{NewDBClient(ctx)}
}

func NewAdminDaoDB(db *gorm.DB) *AdminDao {
	return &AdminDao{db}
}

// 是否存在此管理员
func (dao *AdminDao) ExistOrNotByAdminName(userName string) (admin *model.Admin, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Admin{}).Where("user_name=?", userName).Find(&admin).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return admin, true, nil
}

// 管理员列表
func (dao *AdminDao) ListAdmin(pageNum int, pageSize int) (admins []model.Admin, err error) {
	offset := (pageNum - 1) * pageSize
	err = dao.DB.Model(&model.Admin{}).Offset(offset).Limit(pageSize).Find(&admins).Error
	return
}

//删除管理员
func (dao *AdminDao) DeleteByAdminName(admin *model.Admin) error {
	return dao.DB.Delete(&admin).Error
}
