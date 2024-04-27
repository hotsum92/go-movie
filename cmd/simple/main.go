package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("Server started at port 8080")
	httpServer.Addr = ":8080"
	log.Println(httpServer.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))

	data, err := ioutil.ReadFile("sample.mp4")

	if err != nil {
		log.Println("Error reading file: ", err)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.Write(data)
}
