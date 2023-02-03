package main

import (
	"fmt"
	"giao/practice"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	start := time.Now()
	var count int32 = 0

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(practice.GetUser(strconv.Itoa(i)))
			atomic.AddInt32(&count, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds(), "次", count)
}
