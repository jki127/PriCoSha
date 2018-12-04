package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

//FGD holds Data of Friend Group Page session
type FGD struct {
	Logged           bool
	Username         string
	UserFriendGroups []*b.FriendGroup
}

func friendGroupHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged on and cannot acces friend groups
		log.Println("User was not loggin in and cannot access friend groups.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	CurrentFGD := FGD{
		Logged:           true,
		Username:         cookie.Value,
		UserFriendGroups: b.GetFriendGroup(cookie.Value),
	}
	// username := cookie.Value

	// UserFriendGroups := b.GetFriendGroup(username)

	// t := template.Must(template.ParseFiles("../web/template/friend_groups.html"))
	// t.Execute(w, UserFriendGroups)

	t, err := template.New("").ParseFiles("../web/template/friend_groups.html", "../web/template/base.html")
	if err != nil {
		log.Println("error")
	}
	t.ExecuteTemplate(w, "base", CurrentFGD)
}
