package service

import (
	"go_fabric/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) //强制转换用于路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExist(basePath) {
		CreateFile(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil //用户专属
}

func UploadProductToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId)) //强制转换用于路径拼接
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExist(basePath) {
		CreateFile(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil //用户专属
}

// 判断文件目录是否存在
func DirExist(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 创建文件目录
func CreateFile(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
