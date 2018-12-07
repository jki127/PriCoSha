package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
)

func formDeleteFriendHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	queryData := url.Query()

	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User is not logged in and cannot delete friends.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

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

	cookie, err := r.Cookie("deleteFriendErr")
	var errMsg string
	var isErr bool

	if err == nil {
		isErr = true
		if cookie.Value == "empty" {
			errMsg = "Data fields were left empty. Please retry."
		} else if cookie.Value == "nonexistent" {
			errMsg = "Person you are trying to delete is not in the friend group. Please retry."
		}
	} else {
		isErr = false
	}

	data := struct {
		IsErr      bool
		ErrMsg     string
		FGName     string
		OwnerEmail string
	}{
		isErr,
		errMsg,
		fgName,
		ownerEmail,
	}

	t := template.Must(template.ParseFiles("../web/template/delete_friend.html"))
	t.Execute(w, data)
}