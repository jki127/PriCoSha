package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

func addVoteHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println(`User must be logged in to vote`)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	itemID, _ := strconv.Atoi(r.PostFormValue("itemID"))

	if !b.UserHasAccessToItem(username, itemID) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	choice := r.PostFormValue("choice")

	b.AddVote(username, itemID, choice)

	http.Redirect(w, r, r.Header.Get("referer"), 307)
	return
}
