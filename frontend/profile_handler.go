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
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")

	if err != nil {
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	logged := true
	username := cookie.Value

	CurrentPD := ProfileData{
		Logged:   logged,
		Username: username,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/profile.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentPD)
}
