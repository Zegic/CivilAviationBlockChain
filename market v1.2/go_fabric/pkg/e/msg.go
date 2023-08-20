package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	//user模块错误
	ErrorExistUser:         "用户已存在",
	NullPassword:           "密码不能为空",
	ErrorFailEncryption:    "密码加密失败",
	UserNotFind:            "用户不存在",
	ErrorPassword:          "密码错误",
	ErrorToken:             "token验证失败",
	ErrorCheckTokenTimeOut: "token过期",
	ErrorUploadFail:        "图片上传失败",
	ErrorSendEmail:         "邮件发送失败",
	ErrorNotActive:         "用户未激活",

	//product模块错误
	ErrorProductImgUpload: "商品图片上传失败",
	ProductNotFind:        "商品不存在",
	ErrorBossForProduct:   "商品与商户不匹配",
	ErrorNotSale:          "商品不在售",

	//收藏夹模块
	ErrorFavoriteExist: "收藏夹已存在",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	} else {
		return msg
	}
}
