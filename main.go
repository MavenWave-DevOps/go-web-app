package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"text/template"
)

const port = ":8083"

type TemplateStuff struct {
	BucketUrl  string
	AppVersion string
}

type TemplateConfig struct {
	TemplatePath    string
	BucketConfig    TemplateStuff
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

	bucketUrl := os.Getenv("BUCKET_URL")
	env := os.Getenv("ENV")

	appVersion := fmt.Sprintf("0.0.11-%s", env)

	templateConfig := TemplateConfig{
		TemplatePath: "./templates/index.tpl",
		BucketConfig: TemplateStuff{
			BucketUrl:  bucketUrl,
			AppVersion: appVersion,
		},
		DestinationPath: "./assets/index.html",
	}

	templateConfig.ParseTemplate()

	r := newRouter()
	http.ListenAndServe(port, r)
}
