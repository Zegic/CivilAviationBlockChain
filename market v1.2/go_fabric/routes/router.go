package routes

import (
	"github.com/gin-gonic/gin"
	api "go_fabric/api/v1"
	"go_fabric/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("/GOCODE/go_fabric/static/"))
	// jr.StaticFile("/avatar.jpg", "/GOCODE/go_fabric/static/imags/avatar/avatar.jpg")
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/") //登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)              //用户信息更新
			authed.POST("avatar", api.UpdateAvatar)         //用户上传头像
			authed.POST("user/send-email", api.SendEmail)   //发送验证邮件
			authed.POST("user/valid-email", api.ValidEmail) //绑定邮箱

			//收藏夹模块
			authed.GET("favorite", api.ShowFavorite)
			authed.POST("favorite", api.CreateFavorite)
			authed.POST("favorite-del/:id", api.DeleteFavorite)

			//地址模块
			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.GET("address", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.POST("del-address/:id", api.DeleteAddress)

			//购物车模块
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.POST("del-carts/:id", api.DeleteCart)

			//订单模块
			//authed.POST("order", api.CreateOrder)
			//authed.GET("orders", api.ListOrder)
			//authed.GET("order/:id", api.ShowOrder)
			//authed.PUT("order/:id", api.UpdateOrder)
			//authed.POST("del-order/:id", api.DeleteOrder)
		}

		//管理员操作
		v1.POST("admin/login", api.AdminLogin)
		//获取管理员列表
		v1.GET("admin/list", api.AdminList)
		admin := v1.Group("admin/")
		admin.Use(middleware.JWT())
		{
			admin.GET("user", api.ListUser)             //获取用户列表
			admin.POST("user-del", api.DeleteUser)      //删除用户
			admin.GET("boss", api.ListBoss)             //商户列表
			admin.POST("boss-del", api.DeleteBoss)      //删除商户
			admin.POST("user-passive", api.PassiveUser) //封禁用户
			admin.POST("boss-passive", api.PassiveBoss)
			admin.POST("admin-del", api.DeleteAdmin) //删除管理员
		}

		//商品操作
		//轮播图
		v1.GET("carousels", api.ListCarousel)
		//商品列表
		v1.GET("products", api.ListProduct)
		//搜索商品
		v1.POST("pro-find", api.SearchProduct)
		//展示商品详细信息
		v1.GET("product/:id", api.ShowProduct)
		//获得商品图片
		v1.GET("imgs/:id", api.ListProductImg)
		//获取商品分类
		v1.GET("categories", api.ListCategory)

		//商户登录
		v1.POST("boss/login", api.BossLogin)
		shops := v1.Group("shops/")
		shops.Use(middleware.JWT())
		{
			//发布商品
			shops.POST("product", api.CreateProduct)
			//下架商品
			shops.POST("no-sale", api.NotSale)
		}
	}
	return r
}
