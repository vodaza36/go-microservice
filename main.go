package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Address entity
type Address struct {
	Firstname string
	Lastname  string
	Street    string
	City      string
}

func handlerWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
	defer r.Body.Close()
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recevice request method: %s ", r.Method)
	decoder := json.NewDecoder(r.Body)
	var addr Address
	if err := decoder.Decode(&addr); err != nil {
		http.Error(w, "Invalid request body!", 400)
		return
	}
	log.Printf("Address received: %+v\n", addr)
	defer r.Body.Close()
}

func main() {
	http.HandleFunc("/", handlerWelcome)
	http.HandleFunc("/address", handlePost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
