package main

import (
	"log"
	"net/http"
	"strconv"

	b "pricosha/backend"
)

func unshareHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User is not logged in and cannot add friends.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fgName := r.PostFormValue("fgName")
	ownerEmail := r.PostFormValue("ownerEmail")
	itemID, _ := strconv.Atoi(r.PostFormValue("itemID"))

	valid := b.UserHasRemoveRights(fgName, ownerEmail, username, itemID)

	if !valid {
		log.Println(`User does not have correct privileges to unshare
			Content_Item`)
		http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
		return
	}

	b.UnshareItem(fgName, ownerEmail, itemID)

	http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
	return
}
