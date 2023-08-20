package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(ctx *gin.Context) {
		ctx.String(200, "hellpo_world")
	})

	router.StaticFile("/avatar.jpg", "/gocode/go_fabric/static/imags/avatar/avatar.jpg")

	router.Run(":8080")
}
