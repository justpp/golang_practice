package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//util.Manger()
	//fmt.Println("23")
	//util.UseFunc(util.Closure(1,3))
	//util.ReadFileByOs("./nginx.conf")
	ls := exec.Command("ls", "-a")
	fmt.Println(ls)
	//value := util.ReadIniByBuf("./application.ini", "common", "database.config.dbname")
	//fmt.Println(value)
}
