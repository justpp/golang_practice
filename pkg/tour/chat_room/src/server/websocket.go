package server

import (
	"giao/pkg/tour/chat_room/src/logic"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func wsHandleFunc(w http.ResponseWriter, r *http.Request) {
	// Accept 从客户端接收 WebSocket 握手，并将连接升级到 WebSocket。
	// 如果 Origin 域与主机不同，Accept 将拒绝握手，除非设置了 InsecureSkipVerify 选项（通过第三个参数 AcceptOptions 设置）。
	// 换句话说，默认情况下，它不允许跨源请求。如果发生错误，Accept 将始终写入适当的响应
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatalln("websocket accept err: ", err)
		return
	}
	nickname := r.FormValue("nickname")
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal", nickname)
		_ = wsjson.Write(
			r.Context(),
			conn,
			logic.NewErrMsg("非法昵称；昵称长度2~20"),
		)
		_ = conn.Close(websocket.StatusUnsupportedData, "nickname illegal")
	}

	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已存在", nickname)
		_ = wsjson.Write(
			r.Context(),
			conn,
			logic.NewErrMsg("昵称已存在"),
		)
		_ = conn.Close(websocket.StatusUnsupportedData, "昵称已存在")
	}

}
