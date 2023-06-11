package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("As always this is a simple index handler")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8000")

}
