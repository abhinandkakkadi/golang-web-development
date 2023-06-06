package main

import (
	"fmt"
	"net/http"
)

// here mux.Handle needs a handler so that we know that http.HandlerFunc is a handler. So we type casted normal function to http.HandlerFunc
// So the reason why we are doing this is let's say we want to initialize database before calling the function or if we want to implement middleware and all we can use this method
func messageHandler(s string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s)
	})

}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/abhinand", messageHandler("abhinand"))
	mux.Handle("/athira", messageHandler("athira"))
	http.ListenAndServe(":8080", mux)

}
