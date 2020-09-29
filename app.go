package main

import (
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

func loader(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

func main() {
	http.HandleFunc("/get", search)              // for search (GET) world in list
	http.HandleFunc("/load", loader)             // for load (POST) world list
	log.Fatal(http.ListenAndServe(":8080", nil)) // error log handler
}
