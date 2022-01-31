package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos []todo

func GettodoHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "todos" variable to json
	todoListBytes, err := json.Marshal(todos)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of todos to the response
	w.Write(todoListBytes)
}

func CreatetodoHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of todo
	todo := todo{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the todo from the form info
	todo.Title = r.Form.Get("title")
	todo.Description = r.Form.Get("description")

	// Append our existing list of todos with a new entry
	todos = append(todos, todo)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Maven Wave")
}
