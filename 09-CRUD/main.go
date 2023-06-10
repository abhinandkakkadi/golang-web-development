package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// encoding/json package is used to convert json to struct object and and struct object to json (marshaling and un marshalling)
type Notes struct {
	Title       string `json:"notes"`
	Description string `json:"description"`
}

var noteStore = make(map[string]Notes)
var i = 0

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/addnotes", addNotes).Methods("POST")            // CREATE
	router.HandleFunc("/notes", viewNotes).Methods("GET")               // SEE ALL THE NOTES
	router.HandleFunc("/updatenote/{id}", updateNote).Methods("PATCH")  // UPDATE A NOTE
	router.HandleFunc("/deletenote/{id}", deleteNote).Methods("DELETE") // DElETE

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("listening")
	server.ListenAndServe()

	// http.ListenAndServe(":8080",router)

}

func addNotes(w http.ResponseWriter, r *http.Request) {

	var notes Notes
	err := json.NewDecoder(r.Body).Decode(&notes)
	if err != nil {
		panic(err)
	}
	i++
	k := strconv.Itoa(i)
	noteStore[k] = notes
	res, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(res)

}

func viewNotes(w http.ResponseWriter, r *http.Request) {

	var notes []Notes
	for _, val := range noteStore {
		notes = append(notes, val)
	}

	res, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func updateNote(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)
	id := query["id"]

	var notes Notes
	err := json.NewDecoder(r.Body).Decode(&notes)
	if err != nil {
		panic(err)
	}

	if _, ok := noteStore[id]; ok {

		delete(noteStore, id)
		noteStore[id] = notes

	} else {
		fmt.Println("the id given does not exist")
	}

	res, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := noteStore[id]; ok {
		delete(noteStore, id)
	} else {
		fmt.Println("the id of note not present")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
