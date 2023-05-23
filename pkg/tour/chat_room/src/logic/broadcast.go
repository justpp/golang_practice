package logic

import (
	"log"
	"time"
)

var globalMsgChanLen = 30

type broadcaster struct {
	users map[string]*User

	enterCh chan *User
	leaveCh chan *User
	msgCh   chan *Message

	// 检查用户是否可以进入
	checkUserCh    chan string
	checkUserCanCh chan bool

	// 用chan阻塞 限制获取用户列表
	reqUserListChan chan struct{}
	userListChan    chan []*User
}

var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enterCh: make(chan *User),
	leaveCh: make(chan *User),
	msgCh:   make(chan *Message, globalMsgChanLen),

	checkUserCh:    make(chan string),
	checkUserCanCh: make(chan bool),

	reqUserListChan: make(chan struct{}),
	userListChan:    make(chan []*User),
}

func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enterCh:
			b.users[user.Nickname] = user

		case user := <-b.leaveCh:
			delete(b.users, user.Nickname)

		case msg := <-b.msgCh:
			for _, user := range b.users {
				if user.Uid == msg.User.Uid {
					continue
				}
				user.MsgCh <- msg
			}

		case nickname := <-b.checkUserCh:
			if _, ok := b.users[nickname]; !ok {
				b.checkUserCanCh <- true
			} else {
				b.checkUserCanCh <- false
			}

		case <-b.reqUserListChan:
			userList := make([]*User, 0, len(b.users))
			for _, user := range b.users {
				userList = append(userList, user)
			}
			b.userListChan <- userList
		}
	}
}

func (b *broadcaster) CanEnterRoom(nickname string) bool {
	b.checkUserCh <- nickname

	return <-b.checkUserCanCh
}
func (b *broadcaster) UserEntering(u *User) {
	b.enterCh <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leaveCh <- u
}

func (b *broadcaster) GetUserList() []*User {
	b.reqUserListChan <- struct{}{}
	return <-b.userListChan
}

func (b *broadcaster) Broadcast(msg *Message) {
	if len(b.msgCh) >= globalMsgChanLen {
		log.Println("broadcast msg queue 满了")
	}
	b.msgCh <- msg
}

type Message struct {
	User    *User     `json:"user,omitempty"`
	Type    int       `json:"type,omitempty"`
	Content string    `json:"content,omitempty"`
	MsgTime time.Time `json:"msgTime,omitempty"`
}

const (
	MsgTypeNormal = iota
	MsgWelcome
	MsgTypeUserEnter
	MsgTypeUserLeave
	MsgTypeErr
)

var System = &User{}
