package main

import (
	"fmt"
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
		fmt.Println(err)
		panic(err)
	}

}
