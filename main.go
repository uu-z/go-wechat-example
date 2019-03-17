package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uu-z/gin-wechat/api"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello	World!",
	})
}

func main() {
	app := gin.Default()
	app.GET("/", hello)
	app.GET("/wx", api.WechatIndex)
	app.GET("/wx/redirect", api.WechatRedirect)
	app.GET("/wx/callback", api.WechatCode2token)

	app.Run(":8080")
}
