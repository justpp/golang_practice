package main

import (
	"fmt"
	"giao/jd"
)

func main() {
	j := jd.JDInit()
	err := j.GetQrCode()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = j.CheckScan()
	if err != nil {
		return
	}
}
