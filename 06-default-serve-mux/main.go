package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a handler function")
	})
	http.HandleFunc("/abhinand", messageHandler)
	http.ListenAndServe(":8080", nil)

}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("this is the second handler function")
}
