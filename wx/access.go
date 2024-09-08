package wx

import (
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"net/http"
)

// 使用memcache保存access_token，也可选择redis或自定义cache
var wx = wechat.NewWechat()
var memory = cache.NewMemory()
var cfg = &offConfig.Config{
	AppID:     "",
	AppSecret: "",
	Token:     "",
	// EncodingAESKey: "xxxx",
	Cache: memory,
}

func HandleMessage() {
	// 获取公众号实例
	officialAccount := wx.GetOfficialAccount(cfg)

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// 获取服务端实例
		server := officialAccount.GetServer(req, rw)

		// 设置消息处理器
		server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
			// 演示：回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		})

		// 处理消息接收以及回复
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 发送回复的消息
		server.Send()
	})

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
