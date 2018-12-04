package main

import (
	"html/template"
	"net/http"
	b "pricosha/backend"
)

// Handles requests to upload content item page
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// Check and make sure user is logged in
	cookie, err := r.Cookie("username")
	var username string
	if err != nil {
		// User is not logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	username = cookie.Value
	// Query database for Friend_Group which user belongs to
	FriendGroupData := b.GetUserFriendGroup(username)
	t := template.Must(template.ParseFiles("../web/template/upload.html"))
	t.Execute(w, FriendGroupData)
}
