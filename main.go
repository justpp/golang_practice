package main

import (
	"fmt"
	"giao/utls"
)

func main() {
	//utls.Manger()
	//fmt.Println("23")
	//utls.UseFunc(utls.Closure(1,3))
	//utls.ReadFileByOs("./nginx.conf")
	value := utls.ReadIniByBuf("./application.ini","common","smarty.api.compile_dir")
	fmt.Println(value)
}
