package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

func acceptTagHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	var username string
	if err != nil {
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	username = cookie.Value

	itemID, _ := strconv.Atoi(r.PostFormValue("itemID"))
	tagger := r.PostFormValue("taggerEmail")
	tagged := r.PostFormValue("taggedEmail")

	if username == tagged {
		b.AcceptTag(tagger, tagged, itemID)

	} else {
		log.Println("frontend:	acceptTagHandler():	Tagged User must be logged in to Accept Tag")
	}

	http.Redirect(w, r, "/tag_manager", http.StatusFound)
	return
}
