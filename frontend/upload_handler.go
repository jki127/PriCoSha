package main

import (
	"html/template"
	"net/http"
	b "pricosha/backend"
)

//UPD holds the necessary data for use in the html handlers
type UPD struct {
	FriendGroupData []*b.FriendGroup
	Username        string
	Logged          bool
}

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
	friendGroupData := b.GetUserFriendGroup(username)

	CurrentUPD := UPD{
		FriendGroupData: friendGroupData,
		Username:        username,
		Logged:          true,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/upload.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentUPD)

}
