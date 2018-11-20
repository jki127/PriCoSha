package main

import (
	"html/template"
	"log"
	"net/http"
	b "PriCoSha/backend"
)

// Port that server listens to http requests on (only edit number value)
var httpPort = ":" + "8080"

// Serves favicon file to favicon requests from browser
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../assets/images/favicon_hearts.ico")
}

// Handles requests to root page (referred to as both / and main)
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Checks for requests to non-existent pages
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	/*
		Template format is used as main is planned as templated page,
		as such nil is passed to t.Execute()
	*/
	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // GET - display login page
		t := template.Must(template.ParseFiles("../web/template/login.html"))
		t.Execute(w, nil)
	} else if r.Method == "POST" { // POST - parse form for input data
		email := r.FormValue("email")
		password := r.FormValue("password")
		if email == "" || password == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			if b.ValidateInfo(email, password) {
				log.Println("User logged in successfully with:", email, password)
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				log.Println("User failed to log in with:", email, password)
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	}
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
	// Test connection to database
	if b.TestDB() == nil {
		log.Println("Database connected successfully!")
	} else {
		log.Fatal("Database could not connect.")
	}
	
	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Start server
	log.Println("Frontend spun up!")
	if http.ListenAndServe(httpPort, nil) != nil {
		log.Fatal("ListenAndServe failed.")
	}
}
