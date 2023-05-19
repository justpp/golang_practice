package main

import (
	"fmt"
	"giao/pkg/tour/chat_room/src/server"
	"log"
	"net"
	"net/http"
)

var (
	addr   = ":9999"
	banner = `
    ____               _____
   |     |    |   /\     |
   |     |____|  /  \    | 
   |     |    | /----\   |
   |____ |    |/      \  |

Go 语言编程之旅 —— 一起用 Go 做项目：ChatRoom，start on：%s
`
)

func main() {
	fmt.Printf(banner, addr)

	server.RegisterHandle()

	l, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(err)
	}

	log.Fatalln(http.Serve(l, nil))
}
