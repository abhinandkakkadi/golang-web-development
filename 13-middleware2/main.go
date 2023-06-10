package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is index handler")
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello authenticated user")
}

func middleware1(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 1 before auth")
		next.ServeHTTP(w, r)
		fmt.Println("middleware 1 after auth")
	})

}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("defenitely reached here")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") != "abhi123" {
				fmt.Println("not authenticated")
				return
			} else {
				next.ServeHTTP(w, r)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	router := mux.NewRouter()
	router.Handle("/index", middleware1(auth(http.HandlerFunc(index)))).Methods("GET")
	// when sending request to this route - http://localhost:8080/message?password=abhi123 - use this format
	router.Handle("/message", middleware1(http.HandlerFunc(message))).Methods("GET")
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}
