package v1

import (
	"github.com/gin-gonic/gin"
	"go_fabric/pkg/util"
	"go_fabric/service"
	"net/http"
)

// 管理员登录
func AdminLogin(c *gin.Context) {
	var adminLogin service.AdminService
	if err := c.ShouldBind(&adminLogin); err == nil {
		res := adminLogin.AdminLogin(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

// 获取管理员列表
func AdminList(c *gin.Context) {
	var listAdmin service.UserListService
	if err := c.ShouldBind(&listAdmin); err == nil {
		res := listAdmin.AdminList(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}

}

// 获取用户列表
func ListUser(c *gin.Context) {
	var listUser service.UserListService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listUser); err == nil {
		res := listUser.ListUser(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}

// 删除用户 软删除
func DeleteUser(c *gin.Context) {
	var deleteUser service.AdminDeleteService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteUser); err == nil {
		res := deleteUser.DeleteUser(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

// 封禁用户
func PassiveUser(c *gin.Context) {
	var passiveUser service.AdminDeleteService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&passiveUser); err == nil {
		res := passiveUser.PassiveUser(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

// 删除管理员
func DeleteAdmin(c *gin.Context) {
	var deleteAdmin service.AdminDeleteService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAdmin); err == nil {
		res := deleteAdmin.DeleteAdmin(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

// 查找用户

// 获取商户列表
func ListBoss(c *gin.Context) {
	var listBoss service.UserListService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listBoss); err == nil {
		res := listBoss.ListBoss(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}

// 删除商户
func DeleteBoss(c *gin.Context) {
	var deleteBoss service.AdminDeleteService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteBoss); err == nil {
		res := deleteBoss.DeleteBoss(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

// 封禁商户
func PassiveBoss(c *gin.Context) {
	var passiveBoss service.AdminDeleteService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&passiveBoss); err == nil {
		res := passiveBoss.PassiveBoss(c.Request.Context(), claims.UserName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}
