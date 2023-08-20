package service

import (
	"context"
	"go_fabric/dao"
	"go_fabric/model"
	"go_fabric/pkg/e"
	"go_fabric/pkg/util"
	"go_fabric/serializer"
)

type BossService struct {
	BossName string `json:"boss_name" form:"boss_name"`
	Password string `json:"password" form:"password"`
}

type SaleService struct {
	//productName string `json:"product_name"`
	ProductId uint `json:"product_id" form:"product_id"`
}

// 商户登录
func (service *BossService) Login(ctx context.Context) serializer.Response {
	var boss *model.Boss
	code := e.Success
	bossDao := dao.NewBossDao(ctx)
	//检查用户是否存在
	boss, exist, err := bossDao.ExistOrNotByBossName(service.BossName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "商户不存在，请先注册",
		}
	}
	//验证密码是否正确
	if boss.CheckPassword(service.Password) == false {
		code = e.ErrorPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
	//分发token
	token, err := util.GenerateToken(boss.ID, service.BossName, 0)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.ErrorToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token验证失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BossToken{Boss: serializer.BuildBoss(boss), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// 下架商品
func (service *SaleService) NotSale(ctx context.Context, id uint) serializer.Response {
	var boss *model.Boss
	code := e.Success
	bossDao := dao.NewBossDao(ctx)
	boss, _ = bossDao.GetBossById(id)

	//检查商户是否激活
	if boss.Status == "passive" {
		code = e.ErrorNotActive
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商户没有激活，没有权限",
		}
	}

	// 检查是否存在该商品
	productDao := dao.NewProductDao(ctx)
	product, exist, err := productDao.ExistOrNotById(service.ProductId)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.ProductNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "不存在该商品",
		}
	}

	//检查该商品是否属于该商户
	if product.BossId != boss.ID {
		code = e.ErrorBossForProduct
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "该商户不存在该商品，无权限",
		}
	}

	err = productDao.NotSale(product.ID)
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
		Data:   "商品名:" + product.Name + "已下架",
	}
}
