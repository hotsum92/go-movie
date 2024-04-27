package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("Server started at port 8080")
	httpServer.Addr = ":8080"
	log.Println(httpServer.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
