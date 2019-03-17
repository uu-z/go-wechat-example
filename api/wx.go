package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

var (
	WC *wechat.Wechat
)

func init() {
	cfg := &wechat.Config{
		AppID:          "wx647e72e3b9a5dc41",
		AppSecret:      "1706cce66f576c0cad9894f1ddb7fb9a",
		Token:          "awsl",
		EncodingAESKey: "fJiV25YoS7l8FWngM7pBMmxYxI9x4sl9FGc3cuO40qY",
	}

	WC = wechat.NewWechat(cfg)
}

func WechatIndex(c *gin.Context) {
	// 传入request和responseWriter
	server := WC.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func WechatRedirect(c *gin.Context) {
	oauth := WC.GetOauth()
	err := oauth.Redirect(c.Writer, c.Request, "https://ngrok.awsl.me/wx/callback", "snsapi_userinfo", "123dd123")
	if err != nil {
		fmt.Println(err)
	}
}

func WechatCode2token(c *gin.Context) {

	oauth := WC.GetOauth()
	code := c.Query("code")
	resToken, err := oauth.GetUserAccessToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	userInfo, err := oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userInfo)

	c.JSON(http.StatusOK, gin.H{
		"data": userInfo,
	})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world!",
	})
}
