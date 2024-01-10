package api

import (
	"code-go/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// WsUpGrader 用于将 HTTP请求升级转换为 WebSocket连接
var WsUpGrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨站, 不校验 Origin 请求头
		return true
	},
}

// WsHandleTest 用于处理 WebSocket 连接
func WsHandleTest(c *gin.Context) {
	writer := c.Writer
	req := c.Request
	// 输出客户端请求信息
	core.LOG.Printf("wsEchoHandle: Req Info: %s %s %s\n", req.Method, req.URL, req.Proto)

	// 把 HTTP 请求升级转换为 WebSocket 连接, 并写出 状态行 和 响应头。
	// conn 表示一个 WebSocket 连接, 调用此方法后状态行和响应头已写出, 不能再调用 writer.WriteHeader() 方法。
	conn, err := WsUpGrader.Upgrade(writer, req, nil)
	if err != nil {
		core.LOG.Printf("upgrade error: %v\n", err)
		return
	}
	defer conn.Close()

	// 输出 WebSocket 连接信息
	core.LOG.Printf("wsEchoHandle: websocket info: RemoteAddr=%v, LocalAddr=%v, Subprotocol=%v\n",
		conn.RemoteAddr(), conn.LocalAddr(), conn.Subprotocol())

	for {
		// 读取下一条 (Text/Binary) 数据消息 (接收到 Close 消息或连接异常断开时, 此方法结束阻塞并返回错误)
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			break
		}
		// 打印读取到的数据消息, msgType 的值为 websocket.TextMessage 或 websocket.BinaryMessage
		fmt.Printf("read client(%v) msg: msgType=%d, msg=%s\n",
			conn.RemoteAddr(), msgType, string(msg))

		// conn.ReadMessage() 只能读取 (Text/Binary) 数据消息, 不能读取 (Ping/Pong/Close) 控制消息。
		// 控制消息通过设置对应的 handler 函数处理, 如:
		//     conn.SetPingHandler(), conn.SetPongHandler(), conn.SetCloseHandler()。
		// 控制消息的 handler 函数均有默认值:
		//     PingHandler  默认为自动回复 PongMessage,
		//     PongHandler  默认什么也不做,
		//     CloseHandler 默认把 CloseMessage 发回对方。

		// 把消息写回客户端
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Printf("write error: %v\n", err)
			break
		}
	}

	fmt.Printf("%v OVER\n", conn.RemoteAddr())
}
