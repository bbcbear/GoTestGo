package main

// ?word=foobar' => ["foobar","boofar"]
import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method is supported.", http.StatusNotFound)

		return
	}
}

func loader(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Print(err)
		}
		x := r.Form
		fmt.Print(x)
		return
	}
}

func main() {
	http.HandleFunc("/get", search)              // for search (GET) world in list
	http.HandleFunc("/load", loader)             // for load (POST) world list
	log.Fatal(http.ListenAndServe(":8080", nil)) // error log handler

	x := "OOsvB"
	Library := []string{"peNn", "Teller", "Boovs", "Furs", "fbDtr", "Nnpe"}
	fmt.Println(Anagramm(x, Library))
}

// not used
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

type Words struct {
	ID   int
	Name []string
}

func Anagramm(x string, Library []string) []string {
	var Res []string
	for _, s := range Library {
		if len(x) == len(s) {
			if SortString(strings.ToLower(x)) == SortString(strings.ToLower(s)) {
				Res = append(Res, s)
			}
		}
	}
	return Res
}
