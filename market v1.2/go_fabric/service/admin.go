package service

import (
	"context"
	"go_fabric/dao"
	"go_fabric/model"
	"go_fabric/pkg/e"
	"go_fabric/pkg/util"
	"go_fabric/serializer"
)

type UserListService struct {
	model.BasePage
}

type AdminService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type AdminDeleteService struct {
	UserName   string `json:"user_name" form:"user_name"`
	Password   string `json:"password" form:"password"`
	DeleteName string `json:"delete_name" form:"delete_name"`
}

// 管理员登录
func (service *AdminService) AdminLogin(ctx context.Context) serializer.Response {
	var admin *model.Admin
	code := e.Success
	userDao := dao.NewAdminDao(ctx)
	//检查用户是否存在
	admin, exist, err := userDao.ExistOrNotByAdminName(service.UserName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "管理员不存在",
		}
	}
	//验证密码是否正确
	if admin.CheckAdminPassword(service.Password) == false {
		code = e.ErrorPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
	//分发token
	token, err := util.GenerateToken(admin.ID, service.UserName, 0)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.ErrorToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.AdminToken{
			Admin: admin,
			Token: token,
		},
		Msg: e.GetMsg(code),
	}
}

// 获取用户列表
func (service *UserListService) ListUser(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	//检查管理员是否存在
	_, exist, err := adminDao.ExistOrNotByAdminName(userName)
	if !exist || err != nil {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	userDao := dao.NewUserDao(ctx)
	users, err := userDao.ListUser(service.PageNum, service.PageSize)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildUsers(users), uint(len(users)))
}

// 获取管理员列表
func (service *UserListService) AdminList(ctx context.Context) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	admins, err := adminDao.ListAdmin(service.PageNum, service.PageSize)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(admins, uint(len(admins)))
}

// 删除用户
func (service *AdminDeleteService) DeleteUser(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success
	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
	if !exist {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.DeleteName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，无法删除",
		}
	}
	err = userDao.DeleteByUserName(user)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// PassiveUser封禁用户
func (service *AdminDeleteService) PassiveUser(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success
	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
	if !exist {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.DeleteName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在",
		}
	}

	err = userDao.PassiveByUserName(user.ID)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "用户名：" + user.UserName + "已被封禁",
	}
}

// Passive封禁管理员
//func (service *AdminDeleteService) PassiveAdmin(ctx context.Context, userName string) serializer.Response {
//	adminDao := dao.NewAdminDao(ctx)
//	code := e.Success
//	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
//	if !exist {
//		code = e.UserNotFind
//		return serializer.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//			Data:   "没有该管理员，无权限",
//		}
//	}
//
//	userDao := dao.NewUserDao(ctx)
//	user, exist, err := userDao.ExistOrNotByUserName(service.DeleteName)
//	if !exist || err != nil {
//		util.LoggerObj.Error(err)
//		code = e.UserNotFind
//		return serializer.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//			Data:   "用户不存在",
//		}
//	}
//
//	err = userDao.PassiveByUserName(user.ID)
//	if err != nil {
//		util.LoggerObj.Error(err)
//		code = e.Error
//		return serializer.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//		}
//	}
//
//	return serializer.Response{
//		Status: code,
//		Msg:    e.GetMsg(code),
//		Data:   "用户名：" + user.UserName + "已被封禁",
//	}
//}

// 删除管理员
func (service *AdminDeleteService) DeleteAdmin(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success
	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
	if !exist {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	admin, exist, err := adminDao.ExistOrNotByAdminName(service.DeleteName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，无法删除",
		}
	}
	err = adminDao.DeleteByAdminName(admin)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   admin.UserName + "已被删除",
	}
}

// 获取商户列表
func (service *UserListService) ListBoss(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success

	//检查管理员是否存在
	_, exist, err := adminDao.ExistOrNotByAdminName(userName)
	if !exist || err != nil {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	bossDao := dao.NewBossDao(ctx)
	bosses, err := bossDao.ListBoss(service.PageNum, service.PageSize)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildBosses(bosses), uint(len(bosses)))
}

// 删除商户
func (service *AdminDeleteService) DeleteBoss(ctx context.Context, userName string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success
	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
	if !exist {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	bossDao := dao.NewBossDao(ctx)
	boss, exist, err := bossDao.ExistOrNotByBossName(service.DeleteName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "商户不存在，无法删除",
		}
	}

	err = bossDao.DeleteByBossName(boss)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildBoss(boss),
	}
}

// 封禁商户
func (service *AdminDeleteService) PassiveBoss(ctx context.Context, userName string) interface{} {
	adminDao := dao.NewAdminDao(ctx)
	code := e.Success
	_, exist, _ := adminDao.ExistOrNotByAdminName(userName)
	if !exist {
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "没有该管理员，无权限",
		}
	}

	bossDao := dao.NewBossDao(ctx)
	boss, exist, err := bossDao.ExistOrNotByBossName(service.DeleteName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "商户不存在",
		}
	}

	err = bossDao.PassiveByUserName(boss.ID)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "商户名：" + boss.BossName + "已被封禁",
	}
}
