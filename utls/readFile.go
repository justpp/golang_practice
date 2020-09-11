package utls

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 读取文件
func ReadFileByOs(fileName string) {
	var rows = make([]byte, 4096)
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("读取错误！")
	}
	n, err := file.Read(rows)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(rows[:n]))
}

// 读取配置文件
func ReadIniByBuf(fileName, sectionName, keyName string) string {
	file, err := os.Open(fileName)
	defer file.Close()
	var currSectionName = ""
	if err != nil {
		fmt.Println("读取错误！")
	}
	reader := bufio.NewReader(file)
	for {
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 去掉空格
		lineStr = strings.TrimSpace(lineStr)
		// 忽略空行 注释
		if lineStr == "" || lineStr[0] == ';' || lineStr[0] == '#' {
			continue
		}
		// 取到段名 判断段名是否是想要的
		if lineStr[0] == '[' && lineStr[len(lineStr)-1] == ']' {
			// 定位section_name
			if sectionName == lineStr[1:len(lineStr)-1] && (currSectionName == "" || currSectionName == sectionName) {
				currSectionName = sectionName
			} else {
				// 遍历完了就没有了 跳出
				break
			}
		}
		// 读取键值对
		part := strings.Split(lineStr, "=")
		if len(part) == 2 && strings.TrimSpace(part[0]) == keyName {
			return part[1]
		}
	}
	return "没有想要的键值"+keyName
}
