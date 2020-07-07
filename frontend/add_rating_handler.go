package main

import (
	"log"
	"net/http"
	"strconv"

	b "pricosha/backend"
)

func addRatingHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)
	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID, err := strconv.Atoi(r.FormValue("itemID"))
	if err != nil {
		log.Println("Could not convert itemID from string to int")
		http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
		return
	}

	valid := b.UserHasAccessToItem(username, ID)
	if !valid {
		log.Println("User cannot rate a post they cannot view.")
		http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
		return
	}

	emoji := r.FormValue("emoji")

	b.AddRatingToDB(emoji, username, ID)

	http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
	return
}
