package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func deleteFriendHandler(w http.ResponseWriter, r *http.Request){
	memEmail := r.Form("email")

	url := r.URL
	redirectStr := "/formDeleteFriend?" + url.RawQuery

	if memEmail == "" {
		cookie := http.Cookie{Name: "deleteFriendErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}

	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]

	if ok := b.ValidateBelongFriendGroup(memEmail, fgName, ownerEmail); !ok { ///checks to see if member belongs in friend group to delete
		b.DeleteFriend(memEmail, fgName, ownerEmail)
		log.Println("Deleted Person with email", memEmail)
	}else{
		cookie := http.Cookie{Name: "addFriendErr", Value: "nonexistent"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}
	clearCookie(&w, r, "addFriendErr")
	http.Redirect(w, r, "/friendgroups", http.StatusFound)
	return

}