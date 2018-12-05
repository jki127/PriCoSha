package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

type FGD struct { //friend group data including list of friend groups own, and friend group belongs
	OwnFriendGroups []*b.FriendGroup
	BelongFriendGroups []*b.FriendGroup
}

func friendGroupHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "addFriendErr")

	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged on and cannot acces friend groups
		log.Println("User was not loggin in and cannot access friend groups.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value
	CurrentFGD := FGD{
		OwnFriendGroups:		b.GetFriendGroup(username),
		BelongFriendGroups:	b.GetBelongFriendGroup(username),
	}
	
	t := template.Must(template.ParseFiles("../web/template/friend_groups.html"))
	t.Execute(w, CurrentFGD)
}
