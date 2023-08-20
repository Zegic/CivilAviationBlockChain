package v1

import (
	"github.com/gin-gonic/gin"
	"go_fabric/pkg/util"
	"go_fabric/service"
	"net/http"
)

// 获取商品图片
func ListProductImg(c *gin.Context) {
	var lisProductImgService service.ListProductImg
	if err := c.ShouldBind(&lisProductImgService); err == nil {
		res := lisProductImgService.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}
