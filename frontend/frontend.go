package main

import (
	b "PriCoSha/backend"
	"html/template"
	"log"
	"net/http"
)

// context is used to send data to template files
type context struct {
	Items []*b.ContentItem
}

// Port that server listens to http requests on (only edit number value)
var httpPort = ":" + "8080"

// Serves favicon file to favicon requests from browser
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../assets/images/favicon_hearts.ico")
}

// Handles requests to root page (referred to as both / and main)
func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainCtx := context{Items: b.GetPubContent()}
	// Checks for requests to non-existent pages
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, mainCtx)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../web/template/login.html"))
	t.Execute(w, nil)
}

func validateLoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "" || password == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		if b.ValidateInfo(email, password) {
			log.Println("User logged in successfully with:", email, password)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			log.Println("User failed to log in with:", email, password)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
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
	http.HandleFunc("/validate", validateLoginHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Start server
	log.Println("Frontend spun up!")
	if http.ListenAndServe(httpPort, nil) != nil {
		log.Fatal("ListenAndServe failed.")
	}
}
