package service

import (
	"context"
	"fmt"
	"go_fabric/conf"
	"go_fabric/dao"
	"go_fabric/model"
	"go_fabric/pkg/e"
	"go_fabric/pkg/util"
	"go_fabric/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //加密密钥
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` // 1.绑定邮箱 2.解绑邮箱 3.更改密码
}

type ValidEmailService struct {
}

// 注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 6 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不符",
		}
	}
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "avatar.JPG",
		Status:   model.Active,
		Money:    "1000",
		Vip:      0,
	}
	//密码是否为空
	if service.Password == "" {
		code = e.NullPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		util.LoggerObj.Error(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	//检查用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		util.LoggerObj.Error(err)
		code = e.UserNotFind
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}
	//验证密码是否正确
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
	//分发token
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
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
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// 用户修改信息
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	//寻找用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)

	//检查用户状态
	if user.Status == "passive" {
		code = e.ErrorNotActive
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户没有激活，没有权限",
		}
	}

	//修改昵称nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
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

// 头像更新
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)

	if user.Status == "passive" {
		code = e.ErrorNotActive
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户没有激活，没有权限",
		}
	}

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//保存图片到本地
	path, err := UploadAvatarLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 发送验证邮件
func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice // 绑定邮箱，修改密码 模板通知

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	//检查用户状态
	if user.Status == "passive" {
		code = e.ErrorNotActive
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户没有激活，没有权限",
		}
	}

	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.ErrorToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token //发送者
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "+", address, -1)
	fmt.Println(mailText)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail) //管理员邮箱
	m.SetHeader("To", service.Email)    //验证者邮箱
	m.SetHeader("Subject", "fabric_mall")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		util.LoggerObj.Error(err)
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	//验证token
	if token == "" {
		code = e.InvalidParams

	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			util.LoggerObj.Error(err)
			code = e.ErrorToken

		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorCheckTokenTimeOut

		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// token解析成功，获取用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		util.LoggerObj.Error(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if operationType == 1 { //1.绑定邮箱
		user.Email = email
	} else if operationType == 2 { //2.解绑邮箱
		user.Email = ""
	} else if operationType == 3 { //3.修改密码
		err = user.SetPassword(password)
		if err != nil {
			util.LoggerObj.Error(err)
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
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
