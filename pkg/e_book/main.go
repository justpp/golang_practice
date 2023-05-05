package main

import (
	"giao/pkg/e_book/cmd"
	"log"
)

func main() {
	// url := "https://m2.ddyueshu.com/wapbook/11082821.html"
	// var tool src.EBook
	// tool.Run(url)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
