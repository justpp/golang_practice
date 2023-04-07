package main

import (
	"giao/pkg/e_book/cmd"
	"log"
)

func main() {
	// url := "https://wap.x33xs.com/33xs/392/392417/"
	// var tool src.EBook
	// tool.Run(url)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
