package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

func friendGroupHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged on and cannot acces friend groups
		log.Println("User was not loggin in and cannot access friend groups.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value
	UserFriendGroups := b.GetFriendGroup(username)

	t := template.Must(template.ParseFiles("../web/template/friend_groups.html"))
	t.Execute(w, UserFriendGroups)
}
