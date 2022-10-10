package main

import (
	"log"
	"net/http"
	"os"
  "fmt"

	"github.com/gorilla/mux"
)

//const port = ":8083"

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
  port := fmt.Sprintf(":%s",os.Getenv("PORT"))

	log.SetOutput(os.Stdout)
	log.Printf("Starting web app on %s\n",  port)
  log.Printf("Market region is: %s\n", os.Getenv("MARKET_REGION"))
  log.Printf("Secret Value: %s", os.Getenv("SECRET_DATA"))

	r := newRouter()
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
