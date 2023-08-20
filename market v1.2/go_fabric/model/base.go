package model

// 分页操作

type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
