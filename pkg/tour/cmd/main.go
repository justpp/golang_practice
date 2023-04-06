package main

import (
	cmd "giao/pkg/tour/cmd/repo"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
