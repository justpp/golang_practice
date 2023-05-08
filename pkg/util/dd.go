package util

import (
	"fmt"
	"os"
)

func DD(v ...interface{}) {
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
	os.Exit(1)
}
