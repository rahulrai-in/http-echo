package main

import (
	"net/http"
	"os"
	"time"
)

var prefix string

func main() {
	port := os.Args[1]
	prefix = os.Args[2]
	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":"+port, nil)
}

func echoHandler(writer http.ResponseWriter, request *http.Request) {
	request.Write(writer)
	writer.Write([]byte(prefix + " says okay at " + time.Now().String()))
}
