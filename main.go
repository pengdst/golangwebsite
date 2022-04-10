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

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
