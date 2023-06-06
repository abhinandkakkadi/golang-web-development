package main

import (
	"fmt"
	"net/http"
)

// if you are confused about how the default server is registered with a particular path. It is through http.Handle and http.HandleFunc.

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a handler function")
	})
	http.HandleFunc("/abhinand", messageHandler)
	http.ListenAndServe(":8080", nil) //default serveMux is used for routing

	// here th normal handler function athiraHandler is type casted to http.HandlerFunc, to make it a http.Handler (interface - not function)
	// and then that handler can be used as second parameter to http.Handle function as it is a http.Hanlder type.
	athiHandler := http.HandlerFunc(athiraHandler)
	http.Handle("/athira", athiHandler)

}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("this is the second handler function")
}

func athiraHandler(w http.ResponseWriter, r *http.Request) {

}
