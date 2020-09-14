package util

import (
	"fmt"
	"os"
	"time"
)

func WritSomethingToFile(s string) {
	var dirName = "./log"
	// 判断文件夹是否存在
	dirExists, _ := DirExists(dirName)
	if !dirExists {
		_ = os.Mkdir(dirName, 0644)
	}

	var fileName = dirName + "/" + time.Now().Format("2006-01-02") + ".txt"
	fmt.Println(fileName)

	fileObj, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败！")
	}
	if fileObj != nil {
		_, _ = fileObj.Write([]byte("giaogiaogai"))
	}
}

func DirExists(dirName string) (bool, error) {
	_, err := os.Stat(dirName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, nil
}
