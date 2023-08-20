package dao

import (
	"context"
	"go_fabric/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 是否存在该用户名用户
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

// 添加用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据Id寻找用户
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// 查询所有用户 列表
// offset表示偏移量，pageSize表示每页数据量，pageNum表示当前查询的页码数
func (dao *UserDao) ListUser(pageNum int, pageSize int) (users []model.User, err error) {
	offset := (pageNum - 1) * pageSize
	err = dao.DB.Model(&model.User{}).Offset(offset).Limit(pageSize).Find(&users).Error
	return
}

// 根据id修改用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}

// 根据用户名删除用户
func (dao *UserDao) DeleteByUserName(user *model.User) error {
	return dao.DB.Delete(&user).Error
}

// 封禁用户
func (dao *UserDao) PassiveByUserName(uId uint) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Update("status", "passive").Error
}
