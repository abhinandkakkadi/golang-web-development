package main

import (
	"fmt"
	"net/http"
)

// we will create a custom handler and use it as a handler (the second parameter for the http.Handle function since that function need a handler)

// now this customHandler is a http.Handler since it implements http.Handler interface. and it can be used in places where we use http.Handler
type customHandler struct {
	m string
}

func (c *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c.m)
}

func main() {

	mux := http.NewServeMux()
	c1 := &customHandler{"abhinand"} // now this is a handler as it implements ServeHTTP function
	c2 := &customHandler{"athira"}

	mux.Handle("/abhinand", c1)
	mux.Handle("/athira", c2)
	http.ListenAndServe(":8080", mux)
}
