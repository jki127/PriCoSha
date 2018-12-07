package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func addBioHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User is not logged in and cannot add friends.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	bio := r.FormValue("biography")

	b.AddBioToDB(username, bio)

	http.Redirect(w, r, "/profile", http.StatusFound)
}
