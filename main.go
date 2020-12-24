package main

import (
	"giao/util"
)

func main() {
	//util.Manger()
	//fmt.Println("23")
	//util.UseFunc(util.Closure(1,3))
	//util.ReadFileByOs("./nginx.conf")

	//value := util.ReadIniByBuf("./application.ini", "common", "database.config.dbname")
	//fmt.Println(value)
	//util.WritOS("os")
	//util.WriteBuff("buf")
	//util.WriteIoUtil("ioUtil") // 文件写入

	var str1 = "网站高并发解决方案"
	var str2 = "如何解决网站高并发"

	util.LongestCommSub(str1, str2)

}
