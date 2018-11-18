package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	b "pricosha/backend"
)

// Port that server listens to http requests on (only edit number value)
var httpPort = ":" + "8080"

func main() {
	if b.TestDB() == nil {
		log.Println("Database connected successfully!")
	} else {
		log.Fatal("Database could not be contacted.")
	}

	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/logout", logoutHandler)

	// Start server
	log.Println("Frontend spun up!")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe failed.")
	}
}

// Handles requests to error pages
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		http.ServeFile(w, r, "../web/static/not_found.html")
	default:
		http.ServeFile(w, r, "../web/static/unknown.html")
	}
}

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

	// Currently prints whether user is logged in or not.
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println("User is not logged in.")
	} else {
		log.Println("User is logged in as:", cookie.Value)
	}

	data := b.GetPubContent()
	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, data)
}

// Handles requests to login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/static/login.html")
}

// Handles requests to validate user data and sets cookies accordingly
func validateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	// This check should be revised later
	if username == "" || password == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if ok := b.ValidateInfo(username, password); ok {
		log.Println("User logged in with:", username, password)
		cookie := http.Cookie{Name: "username", Value: username}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		log.Println("User failed to log in with:", username, password)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// Handles requests to logout and deletes cookies accordingly
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		log.Println("User was not logged in and cannot be logged out.")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		cookie := http.Cookie{
			Name:    "username",
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(w, &cookie)
		log.Println("User successfully logged out.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
