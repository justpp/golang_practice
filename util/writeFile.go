package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const dirName = "./log"

func init() {
	// 判断文件夹是否存在
	IsExists, _ := IsExists(dirName)
	if !IsExists {
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

func IsExists(dirName string) (bool, error) {
	_, err := os.Stat(dirName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, nil
}

func CreateFile(fileName string, url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("get file err", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get file err", err)
		return
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("create file err", err)
		return
	}
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Write file err", err)
		return
	}
}
