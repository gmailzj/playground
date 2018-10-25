package main

import (
	"fmt"
	"net/http"
    // "os"
	"github.com/gorilla/mux"
)

func main() {
//     for index, key := range os.Environ() {
// 		fmt.Println("index = ", index, " key = ", key)
// 	}
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8090", r)
}