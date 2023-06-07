package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/home", messageHandler)

	// If we want to configure the server or if we want to customize the behavior of the server we can use Server struct present inside the http package,
	// and setup some functionlities like setup the port, ReadTimeout, WriteTimeout etc.
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok this is message handler")
}
