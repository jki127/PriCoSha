package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

//FGD holds Data of Friend Group Page session
type FGD struct {
	Logged             bool
	Username           string
	OwnFriendGroups    []*b.FriendGroup
	BelongFriendGroups []*b.FriendGroup
}

func friendGroupHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "addFriendErr")

	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged on and cannot access friend groups
		log.Println("User was not logged in and cannot access friend groups.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value

	CurrentFGD := FGD{
		Logged:             true,
		Username:           username,
		OwnFriendGroups:    b.GetFriendGroup(username),
		BelongFriendGroups: b.GetBelongFriendGroup(username),
	}
	t := template.Must(template.New("").ParseFiles("../web/template/friend_groups.html",
		"../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentFGD)
}
