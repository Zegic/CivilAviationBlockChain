package serializer

import (
	"go_fabric/conf"
	"go_fabric/model"
)

// Vo 向前端展示
type UserVo struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) UserVo {
	return UserVo{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		//Type:       ,  用户类型
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []model.User) (users []UserVo) {
	for _, item := range items {
		user := BuildUser(&item)
		users = append(users, user)
	}
	return users
}
