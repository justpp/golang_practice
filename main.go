package main

import (
	"fmt"
	"giao/util"
)

func main() {
	//util.Manger()
	//fmt.Println("23")
	//util.UseFunc(util.Closure(1,3))
	//util.ReadFileByOs("./nginx.conf")
	value := util.ReadIniByBuf("./application.ini", "common", "smarty.api.compile_dir")
	fmt.Println(value)
}
