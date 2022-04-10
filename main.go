package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)

	log.Println("Starting Port 8080")

	err := http.ListenAndServe(":", mux)
	if err != nil {
		panic(err)
	}

}
