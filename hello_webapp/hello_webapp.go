package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr: "localhost:8080",
	}
	log.Fatal(server.ListenAndServe())
}
