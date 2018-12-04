package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

/*
TMD stands for TagManagerData and holds all data necessary
for Tag Manager page to function
*/
type TMD struct {
	Logged          bool // true if logged in, false otherwise
	Username        string
	PendingTagData  []*b.Tag
	AcceptedTagData []*b.Tag
}

func manageTagHandler(w http.ResponseWriter, r *http.Request) {
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

	CurrentTMD := TMD{
		Logged:          logged,
		Username:        username,
		PendingTagData:  b.GetPendingTags(username),
		AcceptedTagData: b.GetAcceptedTags(username),
	}
	t, err := template.New("").ParseFiles("../web/template/tagmanager2.html", "../web/template/base.html")
	if err != nil {
		log.Println("error")
	}
	t.ExecuteTemplate(w, "base", CurrentTMD)
}
