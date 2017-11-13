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
	r.HandleFunc("/{id:[1-9][0-9]*}", todoApp.Delete).Methods("DELETE")
	r.HandleFunc("/{id:[1-9][0-9]*}", todoApp.Update).Methods("PUT")
	//http.Handle("/", r)

	http.ListenAndServe(":5555", r)
}
