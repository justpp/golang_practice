package main

import (
	"giao/pkg/e_book/src"
)

func main() {
	url := "https://wap.x33xs.com/33xs/392/392417/4537075.html"
	// url := "https://wap.x33xs.com/33xs/392/392417/78221204_2.html"

	// url := "https://wap.x33xs.com/33xs/392/392417/100974283.html"
	// url := "https://wap.x33xs.com/33xs/392/392417/100974283_3.html"
	var tool src.EBook

	tool.Run(url)
}
