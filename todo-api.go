package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/MuensterGophers/gruppe1-todoapp/todo"
)

func main() {
	todoApp := todo.NewController()

	r := mux.NewRouter()
	r.HandleFunc("/", todoApp.List).Methods("GET")
	r.HandleFunc("/", todoApp.Create).Methods("POST")
	//http.Handle("/", r)

	http.ListenAndServe(":5555", r)
}
