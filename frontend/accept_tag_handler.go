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

	url := r.URL
	queryData := url.Query()
	itemID, _ := strconv.Atoi(queryData["iid"][0])
	tagger := queryData["ter"][0]
	tagged := queryData["ted"][0]

	if username == queryData["ted"][0] {
		b.AcceptTag(tagger, tagged, itemID)

	} else {
		log.Println("frontend:	acceptTagHandler():	Tagged User must be logged in to Accept Tag")
	}

	http.Redirect(w, r, "/tag_manager", http.StatusFound)
	return
}
