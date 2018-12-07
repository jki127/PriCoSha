package main

import (
	"log"
	"net/http"
	"strconv"

	b "pricosha/backend"
)

func changePrivilegeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println(`frontend: changePrivilegeHandler(): User is not logged in
			and thus cannot change privileges`)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value
	owner := r.PostFormValue("ownerEmail")

	if username != owner {
		log.Println(`frontend: changePrivilegeHandler(): Only the owner of the
			group can change privileges`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	fgName := r.PostFormValue("fgName")
	member := r.PostFormValue("memberEmail")
	action, _ := strconv.Atoi(r.PostFormValue("actionType"))

	b.ChangePrivilege(fgName, owner, member, action)

	http.Redirect(w, r, "/managePrivilege", 307)
	return
}
