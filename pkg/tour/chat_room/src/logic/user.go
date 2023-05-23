package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type User struct {
	Uid      int           `json:"uid,omitempty"`
	Nickname string        `json:"nickname,omitempty"`
	EnterAt  time.Time     `json:"enter_at"`
	Addr     string        `json:"addr,omitempty"`
	Token    string        `json:"token"`
	MsgCh    chan *Message `json:"-"`

	conn  *websocket.Conn
	isNew bool
}

var globalUid uint32 = 0

func (u *User) SendMsg(ctx context.Context) {
	for msg := range u.MsgCh {
		err := wsjson.Write(ctx, u.conn, msg)
		if err != nil {
			fmt.Println("wsjson.Write err", err)
		}
	}
}

func (u *User) CloseMsgCh() {
	close(u.MsgCh)
}

func (u *User) ReceiveMsg(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)

	err = wsjson.Read(ctx, u.conn, &receiveMsg)

	if err != nil {
		//判断连接状态 连接正常关闭不算错误
		if errors.As(err, &websocket.CloseError{}) {
			return nil
		} else if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	msg := NewUserMsg(u, receiveMsg["content"], receiveMsg["send_time"])

	fmt.Println("ReceiveMsg", receiveMsg)
	Broadcaster.Broadcast(msg)
	return nil
}

func NewUser(conn *websocket.Conn, token, nickname, addr string) *User {
	user := &User{
		Nickname: nickname,
		EnterAt:  time.Now(),
		Addr:     addr,
		Token:    token,
		MsgCh:    make(chan *Message, 32),
		conn:     conn,
	}

	// 若有token则检验token
	if user.Token != "" {
		uid, err := parseTokenAndValid(user.Token, user.Nickname)
		if err != nil {
			user.Uid = uid
		}
	}

	// 无token则赋值uid token
	if user.Uid == 0 {
		user.Uid = int(atomic.AddUint32(&globalUid, 1))
		user.Token = getToken(user.Uid, user.Nickname)
		user.isNew = true
	}

	return user

}

func parseTokenAndValid(token, nickname string) (int, error) {

	secret := "-"
	pos := strings.LastIndex(token, "uid")

	messageMac, err := base64.StdEncoding.DecodeString(token[:pos])
	if err != nil {
		return 0, err
	}

	uid, _ := strconv.Atoi(string(token[pos+3]))

	message := fmt.Sprintf("%s%s%d", nickname, secret, uid)

	mac := validateMac([]byte(message), messageMac, []byte(secret))
	if mac {
		return uid, nil
	}
	return 0, errors.New("token is illegal")
}

func getToken(uid int, nickname string) string {
	secret := "-"
	message := fmt.Sprintf("%s%s%d", nickname, secret, uid)
	messageMac := macSha256([]byte(message), []byte(secret))
	return fmt.Sprintf("%suid%d", base64.StdEncoding.EncodeToString(messageMac), uid)
}

func macSha256(message, secret []byte) []byte {
	hash := hmac.New(sha256.New, secret)

	hash.Write(message)
	return hash.Sum(nil)
}

func validateMac(message, messageMac, secret []byte) bool {
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)
	sum := hash.Sum(nil)
	return hmac.Equal(sum, message)
}
