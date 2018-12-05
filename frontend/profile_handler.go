package main

import (
	"html/template"
	"log"
	"net/http"
)

//ProfileData holds info of a Person in the DB i.e. the user
type ProfileData struct {
	Logged   bool
	Username string
	// Fname    string
	// Lname    string
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	var logged bool
	var username string
	if err != nil {
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	logged = true
	username = cookie.Value

	// first, last := b.GetProfileData(username)

	CurrentPD := ProfileData{
		Logged:   logged,
		Username: username,
		// Fname:    first,
		// Lname:    last,
	}
	// t := template.Must(template.ParseFiles("../web/template/profile.html"))
	// t.Execute(w, CurrentPD)

	t := template.Must(template.New("").ParseFiles("../web/template/profile.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentPD)
}
