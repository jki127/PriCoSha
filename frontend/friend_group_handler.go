package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

func friendGroupHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "addFriendErr")
	clearCookie(&w, r, "deleteFriendErr")

	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged on and cannot access friend groups
		log.Println("User was not logged in and cannot access friend groups.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value

	data := struct {
		Logged          bool
		Username        string
		OwnFriendGroups []*b.FriendGroup
		BFGData         []*b.BFGDataElement
	}{
		true,
		username,
		b.GetFriendGroup(username),
		b.GetBelongFriendGroup(username),
	}

	t := template.Must(template.New("").ParseFiles("../web/template/friend_groups.html",
		"../web/template/base.html"))
	t.ExecuteTemplate(w, "base", data)
}
