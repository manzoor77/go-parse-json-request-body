package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func personCreate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p Person

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", p)
	fmt.Fprintf(w, "Person Name: %+v", p.Name)
	fmt.Fprintf(w, "Person Age: %+v", p.Age)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/person/create", personCreate)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
