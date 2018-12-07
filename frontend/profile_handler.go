package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	b "pricosha/backend"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")

	if err != nil {
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	logged := true
	username := cookie.Value

	firstName, lastName, bio, bioBool := b.GetProfileData(username)
	fgd := b.GetFriendGroup(username)
	bfgd := b.GetBelongFriendGroup(username)
	pendingTags := b.GetPendingTags(username)
	// pubCont := b.GetPubContent()
	privCont := b.GetUserContent(username)
	friends := b.GetFriendsList(username)

	CurrentPD := struct {
		Logged             bool
		Username           string
		Fname              string
		Lname              string
		Bio                string
		FriendGroups       []*b.FriendGroup
		BelongFriendGroups []*b.BFGDataElement
		PendingTags        []*b.Tag
		// PublicItems        []*b.ContentItem
		PrivateItems []*b.ContentItem
		FriendsList  []*b.FriendStruct
		HasBio       bool
		UserAvatar   int
	}{
		logged,
		username,
		firstName,
		lastName,
		bio,
		fgd,
		bfgd,
		pendingTags,
		// pubCont,
		privCont,
		friends,
		bioBool,
		rand.Intn(24),
	}

	t := template.Must(template.New("").ParseFiles("../web/template/profile.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", CurrentPD)
}
