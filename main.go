package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8083"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	//Hello world handler
	r.HandleFunc("/hello", Handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/todo", GettodoHandler).Methods("GET")
	r.HandleFunc("/todo", CreatetodoHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(staticFileDirectory)).Methods("GET")
	return r
}

func main() {
	fmt.Printf("Starting web app on %s", port)

	r := newRouter()
	http.ListenAndServe(port, r)
}
