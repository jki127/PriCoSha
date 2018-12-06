package main

import (
	"log"
	"net/http"
	b "pricosha/backend"

)

func chooseAddedFriendHandler(w http.ResponseWriter, r *http.Request){
	url := r.URL
	queryData := url.Query()

	redirectStr := "/duplicateNames?" + url.RawQuery

	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
	userEmail := queryData["ue"][0]

	if ok := b.ValidateBelongFriendGroup(userEmail, fgName, ownerEmail); ok {
		b.AddFriend(userEmail, fgName, ownerEmail)
		log.Println("Added Person with email", userEmail)
	}else{
		cookie := http.Cookie{Name: "addFriendErr", Value: "alreadyBelongs"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}

}