package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func function1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("application handler 1")
}

func function2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("second application handler")
}

func middleware1(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("request method = ",r.Method)
		fmt.Println("before calling 2nd middleware in middleware 1 ",time.Now())
		next.ServeHTTP(w,r)
		fmt.Println("after returning from second middleware in middleware 1",time.Now())
	})

}

func middleware2(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("this is the middleware 2 before calling application handler")
		next.ServeHTTP(w,r)
		fmt.Println("this is in middleware 2 after returning from application handler going to return to middleware 1")
	})
}

func main() {

	handler1 := http.HandlerFunc(function1)
	handler2 := http.HandlerFunc(function2)

	router := mux.NewRouter()
	router.Handle("/abhinand",middleware1(handler1))
	router.Handle("/athira",middleware1(middleware2(handler2)))
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}