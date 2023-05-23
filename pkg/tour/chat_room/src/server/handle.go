package server

import (
	"giao/pkg/tour/chat_room/src/logic"
	"net/http"
	"os"
	"path/filepath"
)

var (
	rootDir = "./"
)

func RegisterHandle() {

	inferRootDir()
	// 广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/ws", wsHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
}

func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(dir string) string
	infer = func(dir string) string {
		if exists(dir + "/src/template") {
			return dir
		}
		return infer(filepath.Dir(dir))
	}
	rootDir = infer(cwd)
}

func exists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}
