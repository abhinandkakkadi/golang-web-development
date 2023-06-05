package main

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi I am Abhinand")

}

func main() {

	mux := http.NewServeMux()

	// type casted a handle function to http.HandlerFunc to use it as a handler
	m := http.HandlerFunc(messageHandler)
	mux.Handle("/", m)
	http.ListenAndServe(":8080", mux)

}
