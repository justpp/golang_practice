package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const dirName = "./log"

func init() {
	// 判断文件夹是否存在
	dirExists, _ := DirExists(dirName)
	if !dirExists {
		_ = os.Mkdir(dirName, 0644)
	}
}

func WritOS(s string) {

	var fileName = dirName + "/" + time.Now().Format("2006-01-02") + ".txt"
	fileObj, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败！")
	}
	if fileObj != nil {

		_, _ = fileObj.Write([]byte(s + "\n"))
		_, _ = fileObj.Write([]byte("aha\n"))
		_, _ = fileObj.Write([]byte("aha\n"))
		_, _ = fileObj.Write([]byte("bcc\n"))
	}
}

func WriteBuff(s string) {
	var fileName = dirName + "/" + time.Now().Format("2006-01-02") + ".txt"
	fileObj, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败！")
	}
	if fileObj != nil {
		defer fileObj.Close()
		fBuf := bufio.NewWriter(fileObj)
		_, _ = fBuf.WriteString(s)
		_ = fBuf.Flush()
	}
}

func WriteIoUtil(s string) {
	var fileName = dirName + "/" + time.Now().Format("2006-01-02") + ".txt"
	err := ioutil.WriteFile(fileName, []byte(s), 0644)
	if err != nil {
		fmt.Println(err)
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
