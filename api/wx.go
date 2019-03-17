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
	server := WC.GetServer(c.Request, c.Writer)
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
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
