package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Boss struct {
	gorm.Model
	BossName       string `gorm:"unique"`
	PasswordDigest string
	ShopName       string //店铺名
	Status         string
	Avatar         string //头像
	Email          string
	Info           string //描述
}

func (boss *Boss) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	boss.PasswordDigest = string(bytes)
	return nil
}

func (boss *Boss) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(boss.PasswordDigest), []byte(password))
	return err == nil
}
