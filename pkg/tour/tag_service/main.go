package main

import (
	"giao/pkg/tour/tag_service/run_server"
	"log"
	"net/http"
)

func main() {
	err := run_server.RunServer("9991")
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("run_server err: %s", err)
	}
}
