package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	b "pricosha/backend"
)

func addCommentHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)
	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID, err := strconv.Atoi(r.FormValue("itemID"))
	if err != nil {
		log.Println("Could not convert itemID from string to int")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	NewComment := b.Comment{
		ItemID:      ID,
		Email:       username,
		Body:        r.FormValue("body"),
		CommentTime: time.Now(),
	}

	b.ExecInsertComment(NewComment)
	log.Println("Adding comment to", r.FormValue("itemID"))

	http.Redirect(w, r, "/item?iid="+r.FormValue("itemID"), http.StatusFound)
	return
}
