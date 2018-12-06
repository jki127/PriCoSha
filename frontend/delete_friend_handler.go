package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func deleteFriendHandler(w http.ResponseWriter, r *http.Request){
	memberEmail := r.FormValue("memberEmail")
	clearCookie(&w, r, "deleteFriendErr")
	url := r.URL
	redirectStr := "/formDeleteFriend?" + url.RawQuery

	/*
	if memberEmail == "" {
		cookie := http.Cookie{Name: "deleteFriendErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	} */
	//clearCookie(&w, r, "deleteFriendErr")
	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
	

	if ok := b.ValidateBelongFriendGroup(memberEmail, fgName, ownerEmail); !ok { ///checks to see if member belongs in friend group to delete
		b.DeleteFriend(memberEmail, fgName, ownerEmail)
		log.Println("Deleted Person with email", memberEmail)
	}else{
		cookie := http.Cookie{Name: "deleteFriendErr", Value: "nonexistent"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}
	clearCookie(&w, r, "deleteFriendErr")
	http.Redirect(w, r, "/friendgroups", http.StatusFound)
	return

}