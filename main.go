package main

import (
	"context"
	"fmt"
	"giao/util"
	"time"
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

	//var str1 = "网站高并发解决方案"
	//var str2 = "如何解决网站高并发网站高并发解决方案"
	//

	// fmt.Println(float64(5) / float64(3))
	//nums1 := []int{1, 3, 5}
	//nums2 := []int{4, 6, 7, 2, 4, 6, 7}
	//fmt.Println(util.FindMedianSortedArrays(nums1, nums2))
	//util.CreateBinary()
	//res := 0
	//for _, i := range nums2 {
	//	res ^= i
	//	println(res)
	//}
	//println("res", res)

	//util.NNTable()
	util.SwitchDemo()
}

func contextTest()  {
	ctx, cancel := context.WithCancel(context.Background())
	go firstCtx(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("stop all sub goroutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func firstCtx(ctx context.Context) {
	go secondCtx(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("first done")
			return
		default:
			fmt.Println("first running")
			time.Sleep(2 * time.Second)
		}
	}
}

func secondCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("second done")
			return
		default:
			fmt.Println("second running")
			time.Sleep(2 * time.Second)
		}
	}
}