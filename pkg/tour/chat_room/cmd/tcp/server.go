package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type User struct {
	ID          int
	Addr        string
	EnterAt     time.Time
	MessageChan chan Message
}

func (u User) String() string {
	return fmt.Sprintf("ID:%v %v", u.ID, u.Addr)
}

type Message struct {
	OwnerId int
	Content string
}

var (
	IDCount      = 0
	enteringChan = make(chan *User)
	leaveChan    = make(chan *User)
	messageChan  = make(chan Message, 8)
)

func main() {
	l, err := net.Listen("tcp4", ":9999")
	if err != nil {
		panic(err)
	}

	go broadcaster()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChan:
			users[user] = struct{}{}
		case user := <-leaveChan:
			delete(users, user)
			close(user.MessageChan)
		case msg := <-messageChan:
			log.Println(msg)
			for user := range users {
				if user.ID == msg.OwnerId {
					continue
				}
				user.MessageChan <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	user := &User{
		ID:          GenUserId(),
		Addr:        conn.RemoteAddr().String(),
		EnterAt:     time.Now(),
		MessageChan: make(chan Message),
	}
	go sendMsg(conn, user.MessageChan)

	user.MessageChan <- Message{
		user.ID,
		"hello " + user.String(),
	}
	enteringChan <- user
	messageChan <- Message{
		user.ID,
		time.Now().Format("2006-01-02 15:04:03") + " 用户：" + strconv.Itoa(user.ID) + " 登录了",
	}

	// 提出不活跃用户
	var userActive = make(chan struct{})
	go func() {

		d := 10 * time.Second
		timer := time.NewTimer(d)

		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChan <- Message{
			user.ID,
			strconv.Itoa(user.ID) + ": " + input.Text(),
		}
		// 活跃
		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)
	}

	// 用户离开
	leaveChan <- user
	messageChan <- Message{
		user.ID,
		time.Now().Format("2006-01-02 15:04:03") + " 用户：" + strconv.Itoa(user.ID) + " 退出了",
	}

}

func sendMsg(conn net.Conn, msgs chan Message) {
	for s := range msgs {
		_, _ = fmt.Fprintln(conn, s.Content)
	}
}

func GenUserId() int {
	IDCount++
	return IDCount
}
