package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type handler func(conn net.Conn)

var handlers = make(map[string]handler)

func get(url string, handler handler) {
	handlers[url] = handler
}

func resp(conn net.Conn, body []byte) {
	content := fmt.Sprintf("HTTP/2.0 200 OK\r\n\r\n%s", string(body))
	conn.Write([]byte(content))
}

func Server() {
	get("/profile", func(conn net.Conn) {
		resp(conn, []byte("page profile"))
	})
	get("/home", func(conn net.Conn) {
		resp(conn, []byte("hello world"))
	})
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("message", message)
		if strings.HasPrefix(message, "GET") {
			sp := strings.Split(message, " ")
			if handler, ok := handlers[sp[1]]; ok {
				handler(conn)
				conn.Close()
			}
		}
	}
}
