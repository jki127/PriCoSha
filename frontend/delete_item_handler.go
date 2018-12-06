package main

import (
	"net/http"
	"log"
	"strconv"

	b "pricosha/backend"
)

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		// User is not logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID, err := strconv.ParseInt(r.FormValue("itemID"), 10, 64)
	if err != nil {
		log.Println("Could not convert itemID from string to int64")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	b.ExecDeleteContentItem(ID)
	log.Println("Deleting content item:",r.FormValue("itemID"))

	http.Redirect(w, r, "/", http.StatusFound)
	return
}