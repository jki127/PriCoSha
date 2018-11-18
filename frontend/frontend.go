package main

import (
	"html/template"
	"log"
	"net/http"

	"pricosha/backend"
)

// Port that server listens to http requests on (only edit number value)
var httpPort = ":" + "8080"

// Serves favicon file to favicon requests from browser
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../assets/images/favicon_hearts.ico")
}

// Handles requests to root page (referred to as both / and main)
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Checks for requests to non-existant pages
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	data := backend.GetPubContent()
	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, data)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		http.ServeFile(w, r, "../web/static/not_found.html")
	default:
		http.ServeFile(w, r, "../web/static/unknown.html")
	}
}

func main() {
	if backend.TestDB() == nil {
		log.Println("Database connected successfully!")
	} else {
		log.Fatal("Database could not be contacted.")
	}

	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Start server
	log.Println("Frontend spun up!")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe failed.")
	}
}
