package middleware

import (
	"github.com/gin-gonic/gin"
	"go_fabric/pkg/e"
	"go_fabric/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorCheckTokenTimeOut
			}
		}
		if code != e.Success {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
