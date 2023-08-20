package service

import (
	"context"
	"go_fabric/dao"
	"go_fabric/pkg/e"
	"go_fabric/pkg/util"
	"go_fabric/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	categories, err := categoryDao.ListCategory()
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(categories), uint(len(categories)))
}
