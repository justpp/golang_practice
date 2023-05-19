package logic

import (
	"net"
	"time"
)

var msgQueueLen = 10

var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enterCh: make(chan *User),
	leaveCh: make(chan *User),
	msgCh:   make(chan *Message, msgQueueLen),

	checkUserCh:    make(chan string),
	checkUserCanCh: make(chan bool),
}

type broadcaster struct {
	users map[string]*User

	enterCh chan *User
	leaveCh chan *User
	msgCh   chan *Message

	// 检查用户是否可以进入
	checkUserCh    chan string
	checkUserCanCh chan bool
}

func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enterCh:
			b.users[user.Nickname] = user

			// send list

		case user := <-b.leaveCh:
			delete(b.users, user.Nickname)

		// send list

		case msg := <-b.msgCh:
			for _, user := range b.users {
				if user.Uid == msg.User.Uid {
					continue
				}
				user.msgCh <- msg
			}

		case nickname := <-b.checkUserCh:
			if _, ok := b.users[nickname]; !ok {
				b.checkUserCanCh <- true
			} else {
				b.checkUserCanCh <- false
			}
		}
	}
}

func (b *broadcaster) CanEnterRoom(nickname string) bool {
	_, ok := b.users[nickname]
	return ok
}

func (b *broadcaster) sendUserList() {

}

type User struct {
	Uid       int       `json:"uid,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	EnterTime time.Time `json:"enter_time"`
	Addr      string    `json:"addr,omitempty"`

	msgCh chan *Message
	conn  *net.Conn
}

type Message struct {
	User    *User            `json:"user,omitempty"`
	Type    int              `json:"type,omitempty"`
	Content string           `json:"content,omitempty"`
	MsgTime time.Time        `json:"msgTime"`
	Users   map[string]*User `json:"users,omitempty"`
}

const (
	MsgTypeNormal = iota
	MsgTypeSystem
	MsgTypeErr
	MsgTypeUserList
)

var System = &User{}
