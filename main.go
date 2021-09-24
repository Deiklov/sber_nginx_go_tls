package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/mef", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, perec %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root")
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}
