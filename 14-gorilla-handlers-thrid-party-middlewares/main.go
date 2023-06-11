package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the index handler")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the home handler")
}

func main() {

	router := mux.NewRouter()
	// here handlers.LoggingHandler is a middleware which execute items after the application handler is called. i.e after the execution of request.
	// not that middleware can have 2 arguments. it's just that the return type must be a handler
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(index)))
	router.Handle("/home", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(home)))
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
