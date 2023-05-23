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
	// 1. 新用户进来，构建该用户的实例
	token := r.FormValue("token")
	nickname := r.FormValue("nickname")

	if l := len(nickname); l < 2 || l > 20 {
		_ = wsjson.Write(
			r.Context(),
			conn,
			logic.NewErrMsg("非法昵称；昵称长度2~20"),
		)
		_ = conn.Close(websocket.StatusUnsupportedData, "nickname illegal")
	}
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		_ = wsjson.Write(
			r.Context(),
			conn,
			logic.NewErrMsg("昵称已存在"),
		)
		_ = conn.Close(websocket.StatusUnsupportedData, "昵称已存在")
	}

	userHasToken := logic.NewUser(conn, token, nickname, r.RemoteAddr)

	// 2.开启给用户发送消息的 goroutine
	go userHasToken.SendMsg(r.Context())

	// 3. 给当前用户发送欢迎信息
	userHasToken.MsgCh <- logic.NewWelcomeMsg(userHasToken)

	// 避免 token 泄露
	tmpUser := *userHasToken
	user := &tmpUser
	user.Token = ""

	// 给所有用户告知新用户到来
	msg := logic.NewUserEnterMsg(user)
	logic.Broadcaster.Broadcast(msg)

	// 4. 将该用户加入广播器的用列表中
	logic.Broadcaster.UserEntering(user)
	log.Println("user:", nickname, "joins chat")

	// 5. 接收用户消息
	err = user.ReceiveMsg(r.Context())

	// 6. 用户离开
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewUserLeaveMsg(user)
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", nickname, "leaves chat")

	// 根据读取时的错误执行不同的 Close
	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}

}
