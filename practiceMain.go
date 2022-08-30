package main

import (
	"fmt"
	"giao/util"
)

func main() {
	uuid := util.NewUuidGenerator("gg")

	fmt.Println(uuid.Get())
	for i := 0; i < 50; i++ {
		fmt.Println(uuid.GetUint32())
	}
}
