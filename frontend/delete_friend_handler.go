package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func deleteFriendHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User is not logged in and cannot remove friends.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	url := r.URL
	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
	role := b.GetRole(fgName, ownerEmail, username)

	switch role {
	case 0:
		// do nothing
	case 1:
		// do nothing
	default:
		log.Println(`User does not have correct privileges to delete friends.`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	memberEmail := r.FormValue("memberEmail")
	redirectStr := r.Header.Get("referer")

	if memberEmail == "" {
		cookie := http.Cookie{Name: "deleteFriendErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}

	if ok := b.ValidateBelongFriendGroup(memberEmail, fgName, ownerEmail); !ok { ///checks to see if member belongs in friend group to delete
		b.DeleteFriend(memberEmail, fgName, ownerEmail)
		b.RemoveInvalidTags(memberEmail)
		log.Println("Deleted Person with email", memberEmail)
	} else {
		cookie := http.Cookie{Name: "deleteFriendErr", Value: "nonexistent"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}
	clearCookie(&w, r, "deleteFriendErr")

	log.Println(memberEmail)
	http.Redirect(w, r, "/friendgroups", http.StatusFound)
	return
}
