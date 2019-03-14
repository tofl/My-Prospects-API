package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login/{identifier}", login).Methods("GET")

	if err := http.ListenAndServe(":8888", router); err != nil {
		fmt.Println(err)
	}
}

func respond(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func login(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	identifier := params["identifier"]

	if identifier != "" {
		fmt.Println(identifier)
	}


	respond(w, http.StatusOK)
	return
}