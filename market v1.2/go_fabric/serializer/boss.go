package serializer

import (
	"go_fabric/conf"
	"go_fabric/model"
)

// Vo 向前端展示
type BossVo struct {
	ID       uint   `json:"id"`
	BossName string `json:"boss_name"`
	ShopName string `json:"shop_name"`
	Status   string `json:"status"`
	Info     string `json:"info"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildBoss(boss *model.Boss) BossVo {
	return BossVo{
		ID:       boss.ID,
		BossName: boss.BossName,
		ShopName: boss.ShopName,
		Email:    boss.Email,
		Status:   boss.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + boss.Avatar,
		CreateAt: boss.CreatedAt.Unix(),
	}
}

func BuildBosses(items []model.Boss) (bosses []BossVo) {
	for _, item := range items {
		boss := BuildBoss(&item)
		bosses = append(bosses, boss)
	}
	return bosses
}
