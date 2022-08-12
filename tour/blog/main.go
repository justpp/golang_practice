package main

import (
	"giao/tour/blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()

	r := &http.Server{
		Addr:           ":9999",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := r.ListenAndServe()
	if err != nil {
		return
	}
}
