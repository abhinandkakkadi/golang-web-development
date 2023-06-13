package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)



func authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token := r.Header.Get("mytoken")
	fmt.Println(token)
	if token == "abhinand" {
		context.Set(r,"user","abhinand")
		next.ServeHTTP(w,r)
	} else {
		http.Error(w,"not authorized",http.StatusUnauthorized)
	}
} 

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is index handler")
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/",index)

	n := negroni.Classic()
	// actually here I am adding a middleware which will be applicable for all routes in the
	// the above multiplexer

	n.Use(negroni.HandlerFunc(authorize))

	// here we can add all the multiplexers/router to which the above middleware should be applied
	n.UseHandler(mux)
	n.Run(":8080")
}