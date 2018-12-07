package main

import (
	"log"
	"net/http"

	b "pricosha/backend"
)

func changeOwnerHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println(`frontend: changeOwnerHandler(): User is not logged in
			and thus cannot rename the group`)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value
	owner := r.PostFormValue("ownerEmail")

	if username != owner {
		log.Println(`frontend: changeOwnerHandler(): Only the owner of the
			group can rename the group`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	fgName := r.PostFormValue("fgName")
	newOwner := r.PostFormValue("newOwner")

	b.SwapOwner(fgName, owner, newOwner)

	http.Redirect(w, r, "/friendgroups", http.StatusFound)
	return
}
