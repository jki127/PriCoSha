package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

//ProfileData holds info of a Person in the DB i.e. the user
type ProfileData struct {
	Logged             bool
	Username           string
	Fname              string
	Lname              string
	FriendGroups       []*b.FriendGroup
	BelongFriendGroups []*b.FriendGroup
	PendingTags        []*b.Tag
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")

	if err != nil {
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	logged := true
	username := cookie.Value

	firstName, lastName := b.GetProfileData(username)
	fgd := b.GetFriendGroup(username)
	bfgd := b.GetBelongFriendGroup(username)
	pendingTags := b.GetPendingTags(username)

	CurrentPD := ProfileData{
		Logged:             logged,
		Username:           username,
		Fname:              firstName,
		Lname:              lastName,
		FriendGroups:       fgd,
		BelongFriendGroups: bfgd,
		PendingTags:        pendingTags,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/profile.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentPD)
}
