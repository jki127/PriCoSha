package main

import (
	"html/template"
	"net/http"
	b "pricosha/backend"
)

/*
MPD stands for MainPageData and holds all data necessary
for main page to function
*/
type MPD struct {
	Logged   bool // true if logged in, false otherwise
	Username string
	PubData  []*b.ContentItem
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

	t := template.Must(template.New("").ParseFiles("../web/template/main.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentMPD)
}
