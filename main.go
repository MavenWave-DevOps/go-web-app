package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"github.com/gorilla/mux"
)

const port = ":8083"

type TemplateStuff struct {
	BucketUrl string
}

type TemplateConfig struct {
	TemplatePath string
	BucketConfig TemplateStuff
	DestinationPath string
}

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

func (t TemplateConfig) ParseTemplate() {
	tpl, err := template.ParseFiles(t.TemplatePath)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(t.DestinationPath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = tpl.Execute(f, t.BucketConfig)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	f.Close()

}

func main() {
	fmt.Printf("Starting web app on %s", port)

	templateConfig := TemplateConfig{
		TemplatePath: "./templates/index.tpl",
		BucketConfig: TemplateStuff{
			BucketUrl: "https://storage.googleapis.com/public-assets-poc/argo.png",
		},
		DestinationPath: "./assets/index.html",
	}

	templateConfig.ParseTemplate()

	r := newRouter()
	http.ListenAndServe(port, r)
}
