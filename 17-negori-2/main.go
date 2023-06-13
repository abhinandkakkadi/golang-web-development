package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

// func main() {
// 	router := mux.NewRouter()
// 	adminRoutes := mux.NewRouter()
// 	router.Handle("/admin", negroni.New(
//   Middleware1,
//   Middleware2,
//   negroni.Wrap(adminRoutes),
// ))
// }

// the important thing to note here is that we can create multiple router
// so that we can divide between different routes - eg : admin routes and user routes


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",index)
	mux.HandleFunc("/message",message)
	n := negroni.Classic()

	// when writing middleware like this the order really matter - since the middleware1 is added first it will be called first followed my middleware 2 and application handler
	n.Use(negroni.HandlerFunc(middleware1))
	n.Use(negroni.HandlerFunc(middleware2))
	n.UseHandler(mux)
	
	n.Run(":8080")
}


func middleware1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("middleware 1 before calling middleware 2")
	next.ServeHTTP(w,r)
	fmt.Println("middleware 1 after calling middleware 2")
}

func middleware2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("middleware 2 before calling application")
	next.ServeHTTP(w,r)
	fmt.Println("middleware 2 after calling application handler")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")
}


func message(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message handler")
}