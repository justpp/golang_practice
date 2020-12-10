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

	//value := util.ReadIniByBuf("./application.ini", "common", "database.config.dbname")
	//fmt.Println(value)
	//util.WritOS("os")
	//util.WriteBuff("buf")
	//util.WriteIoUtil("ioUtil")
	l1 := &util.ListNode{
		Val: 2,
		Next: &util.ListNode{
			Val: 4,
			Next: &util.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &util.ListNode{
		Val: 5,
		Next: &util.ListNode{
			Val: 6,
			Next: &util.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	//util.ShowListNode(l1)
	//util.ShowListNode(l2)
	res := util.TowSumList(l1, l2)
	fmt.Println("result:")
	util.ShowListNode(res)
}
