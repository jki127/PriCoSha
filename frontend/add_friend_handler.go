package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

func addFriendHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User is not logged in and cannot add friends.")
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
		log.Println(`User does not have correct privileges to add friends.`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	redirectStr := r.Header.Get("referer")
	redirectStr2 := "/duplicateNames?" + url.RawQuery + "&fname=" + fname + "&lname=" + lname

	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

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
		http.Redirect(w, r, redirectStr2, http.StatusFound)
		return
	} else if len(EmailList) == 0 {
		cookie := http.Cookie{Name: "addFriendErr", Value: "nonexistent"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, redirectStr, http.StatusFound)
		return
	}
	clearCookie(&w, r, "addFriendErr")

	userEmail = *EmailList[0]

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
