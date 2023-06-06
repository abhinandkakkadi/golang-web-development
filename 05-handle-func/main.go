package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	// actually here internally message handler handler function is converted to http.HandlerFunc internally to make it http.Handler type so
	// that it can be used as a handler
	mux.HandleFunc("/", messageHandler)
	http.ListenAndServe(":8080", mux)

}

func messageHandler(w http.ResponseWriter, e *http.Request) {

	fmt.Println("this is the message handler")

}
