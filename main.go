package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.NumCPU()
	fmt.Println(n)
}
