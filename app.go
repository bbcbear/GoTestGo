package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is supported.", http.StatusNotFound)
		return
	}
}

func loader(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		r.ParseForm()
		fmt.Fprintln(w, r.Form)
		return
	}
}

func main() {
	http.HandleFunc("/get", search)              // for search (GET) world in list
	http.HandleFunc("/load", loader)             // for load (POST) world list
	log.Fatal(http.ListenAndServe(":8080", nil)) // error log handler
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

type Words struct {
	ID   int
	Name []string
}
