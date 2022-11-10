package practice

import (
	"fmt"
	"time"
)

func WorkerPoolPractice() {
	// 起始时间
	nowTime := time.Now()
	members := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			members <- i + 1
			fmt.Printf("send %v\n", i+1)
		}
		close(members)
	}()
	for member := range members {
		time.Sleep(time.Second * 1)
		fmt.Printf("read %v\n", member)
	}
	// 消耗时间
	endTime := time.Now().Sub(nowTime)
	fmt.Printf("total time token %v\n", endTime.Seconds())
}
