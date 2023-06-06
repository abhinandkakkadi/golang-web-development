package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	// here we are using a separate router mux as it provide some additional functionality.
	// Here we can match a request with path,request method,query parameters, and so on.
	r := mux.NewRouter()
	r.HandleFunc("/abhi",messageHandler).Methods("GET")
	r.HandleFunc("/abhi/{id}",secondHandler).Methods("GET")
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	server.ListenAndServe()
	
}


func messageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("this is the message handler")

}

func secondHandler(w http.ResponseWriter,r *http.Request) {

	query := r.URL.Query()
	id := query.Get("id")
	fmt.Println(id)
	
}