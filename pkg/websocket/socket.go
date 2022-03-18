package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Callback func(text string, ws *websocket.Conn, messageType int) error

func EchoMessage(c *gin.Context, callback Callback) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		//err = ws.WriteMessage(mt, message)
		//if err != nil {
		//	break
		//}
		call := callback(string(message), ws, mt)
		if call != nil {
			break
		}
	}
}

// WsWrite 文件写入
func WsWrite(text string, ws *websocket.Conn) error {
	return ws.WriteMessage(1, []byte(text))
}

/*
bindAddress := "localhost:2303"
r := gin.Default()
r.GET("/xxxx", EchoMessage)
r.Run(bindAddress)
*/
