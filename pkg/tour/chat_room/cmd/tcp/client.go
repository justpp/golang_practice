package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	talk()
}

func talk() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error)

	for {
		conn := connect(errChan)

		go mustCopy(conn, os.Stdin, errChan)

		select {
		case <-sig:
			log.Println("手动退出")
			conn.Close()
			return
		case err := <-errChan:
			log.Println(err)
			conn.Close()
			continue
		}
	}
}

func connect(errChan chan error) net.Conn {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			errChan <- err // 发送EOF到通道
			return
		}
		log.Println("被服务端踢掉了。。。")
		conn.Close()
	}()
	return conn
}

func mustCopy(dst io.Writer, src io.Reader, errChan chan error) {
	_, err := io.Copy(dst, src)
	if err != nil {
		errChan <- errors.New("尝试重新连接") // 发送连接断开的错误
		return
	}
	errChan <- errors.New("发送程序异常的错误")
	return
}
