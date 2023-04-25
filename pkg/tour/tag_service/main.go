package main

import (
	"giao/src/tour/tag_service/run_server"
	"log"
)

func main() {
	err := run_server.RunServer("9991")
	if err != nil {
		log.Fatalf("err:%s", err)
	}
}
