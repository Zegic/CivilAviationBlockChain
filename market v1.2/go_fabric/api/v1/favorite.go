package v1

import (
	"github.com/gin-gonic/gin"
	"go_fabric/pkg/util"
	"go_fabric/service"
	"net/http"
)

func CreateFavorite(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Create(c.Request.Context(),claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}

func ShowFavorite(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&showFavoriteService); err == nil {
		res := showFavoriteService.Show(c.Request.Context(),claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	deleteFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(),claim.ID,c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LoggerObj.Error(err)
	}
}