// This is a placeholder frontend .go file
package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp map[string]int

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../views/main.html"))
	t.Execute(w, temp)
}

func main() {
	http.HandleFunc("/", mainHandler)

	log.Println("Frontend spun up!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
