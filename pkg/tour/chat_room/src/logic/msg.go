package logic

import "time"

func NewErrMsg(err string) *Message {
	return &Message{
		User:    System,
		Type:    MsgTypeErr,
		Content: err,
		MsgTime: time.Now(),
	}
}

func NewUserMsg(user *User, content, clientTime string) *Message {
	msg := &Message{
		User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}

	return msg

}
func NewWelcomeMsg(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgWelcome,
		Content: user.Nickname + "欢迎加入聊天室",
		MsgTime: time.Now(),
	}
}
func NewUserEnterMsg(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserEnter,
		Content: user.Nickname + "加入了聊天室",
		MsgTime: time.Now(),
	}
}
func NewUserLeaveMsg(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserLeave,
		Content: user.Nickname + "离开了聊天室",
		MsgTime: time.Now(),
	}
}
