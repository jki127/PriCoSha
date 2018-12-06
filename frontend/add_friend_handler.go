package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func addFriendHandler(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	url := r.URL

	redirectStr := r.Header.Get("referer")

	if fname == "" || lname == "" {
		cookie := http.Cookie{Name: "addFriendErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}

	// Get Emails of Person's with inputted name
	EmailList := b.GetEmail(fname, lname)
	var userEmail string

	if len(EmailList) > 1 {
		cookie := http.Cookie{Name: "addFriendErr", Value: "duplicates"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	} else if len(EmailList) == 0 {
		cookie := http.Cookie{Name: "addFriendErr", Value: "nonexistent"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}
	clearCookie(&w, r, "addFriendErr")

	userEmail = *EmailList[0]

	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]

	if ok := b.ValidateBelongFriendGroup(userEmail, fgName, ownerEmail); ok {
		b.AddFriend(userEmail, fgName, ownerEmail)
		log.Println("Added Person with email", userEmail)
	} else {
		cookie := http.Cookie{Name: "addFriendErr", Value: "alreadyBelongs"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}

	http.Redirect(w, r, "/friendgroups", http.StatusFound)
	return
}
