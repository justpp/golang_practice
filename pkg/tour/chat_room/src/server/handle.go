package server

import (
	"net/http"
	"os"
	"path/filepath"
)

var (
	rootDir = "./"
)

func RegisterHandle() {

	inferRootDir()
	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/ws", wsHandleFunc)
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
