package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	b "pricosha/backend"
)

// Port that server listens 
var httpPort = ":" + "8080"


type MainPage struct {
	LoggedOn   bool // true if logged in, false otherwise
	Username string
	PubData  []*b.ContentItem
}

func main() {
	if b.TestDB() == nil {
		log.Println("Database connected successfully!")
	} else {
		log.Fatal("Database could not be contacted.")
	}

	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
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

// error pages
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		http.ServeFile(w, r, "../web/static/not_found.html")
	default:
		http.ServeFile(w, r, "../web/static/unknown.html")
	}
}

// Handles requests to root page (referred to as both / and main)
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Checks for requests to non-existant pages
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	cookie, err := r.Cookie("username")
	var logged bool
	var username string
	if err != nil {
		logged = false
		username = ""
	} else {
		logged = true
		username = cookie.Value
	}

	CurrentMP := MainPage{
		LoggedOn:   logged,
		Username: username,
		PubData:  b.GetPubContent(),
	}

	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, CurrentMP)
}

// Handles requests to login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err == nil {
		// User is already logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Check if there was a previous login error
	cookie, err := r.Cookie("logErr")
	var errMsg string
	var isErr bool
	if err == nil {
		isErr = true
		if cookie.Value == "empty" {
			errMsg = "Data fields were left empty. Please retry."
		} else if cookie.Value == "fail" {
			errMsg = "Incorrect login data. Please retry."
		}
	} else {
		isErr = false
	}

	data := struct {
		IsErr  bool
		ErrMsg string
	}{
		isErr,
		errMsg,
	}

	t := template.Must(template.ParseFiles("../web/template/login.html"))
	t.Execute(w, data)
}

