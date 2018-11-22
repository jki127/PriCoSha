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

/*
MPD stands for MainPageData and holds all data necessary
for main page to function
*/
type MPD struct {
	Logged   bool // true if logged in, false otherwise
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
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateLoginHandler)
	http.HandleFunc("/logout", logoutHandler)

	// Start server
	log.Println("Frontend spun up!")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("frontend: main(): ListenAndServe() failed.")
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
	// Checks for requests to non-existent pages
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	// Currently prints whether user is logged in or not.
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

	CurrentMPD := MPD{
		Logged:   logged,
		Username: username,
		PubData:  b.GetPubContent(),
	}

	t := template.Must(template.ParseFiles("../web/template/main.html"))
	t.Execute(w, CurrentMPD)
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

// Handles requests to validate user data and sets cookies accordingly
func validateLoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err == nil {
		// User is already logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	// This check should be revised later
	if username == "" || password == "" {
		cookie := http.Cookie{Name: "logErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if ok := b.ValidateInfo(username, password); ok {
		log.Println("User logged in with:", username, password)
		// Set cookie with user info
		cookie := http.Cookie{Name: "username", Value: username}
		http.SetCookie(w, &cookie)
		// Delete logErr cookie if it exists
		_, err := r.Cookie("logErr")
		if err == nil {
			c := http.Cookie{
				Name:    "logErr",
				Value:   "",
				Expires: time.Unix(0, 0),
			}
			http.SetCookie(w, &c)
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	log.Println("User failed to log in with:", username, password)
	cookie := http.Cookie{Name: "logErr", Value: "fail"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
	return
}

// Handles requests to logout and deletes cookies accordingly
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		log.Println("User was not logged in and cannot be logged out.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	cookie := http.Cookie{
		Name:    "username",
		Value:   "",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
