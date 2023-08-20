package v1

import (
	"github.com/gin-gonic/gin"
	"go_fabric/pkg/util"
	"go_fabric/service"
	"net/http"
)

func BossLogin(c *gin.Context) {
	var bossLogin service.BossService
	if err := c.ShouldBind(&bossLogin); err == nil {
		res := bossLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err)) //400参数错误
		util.LoggerObj.Error(err)
	}
}

func NotSale(c *gin.Context) {
	var NotSaleProductService service.SaleService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&NotSaleProductService); err == nil {
		res := NotSaleProductService.NotSale(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}
