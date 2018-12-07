package main

import (
	"html/template"
	"log"
	"net/http"

	b "pricosha/backend"
)

func managePrivilegesHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "chngPrivErr")

	logged, username := getUserSessionInfo(r)

	if !logged {
		log.Println("User was not logged in and cannot manage privileges.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	owner := r.PostFormValue("ownerEmail")
	group := r.PostFormValue("fgName")
	role := b.GetRole(group, owner, username)

	switch role {
	case 0:
		// do nothing
	case 1:
		// do nothing
	default:
		log.Println(`User does not have correct privileges to manage
			privileges.`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	data := struct {
		Logged     bool
		FGName     string
		OwnerEmail string
		Role       int
		Mods       []*string
		Members    []*string
	}{
		true,
		group,
		owner,
		role,
		b.GetAtRole(group, owner, 1),
		b.GetAtRole(group, owner, 2),
	}

	t := template.Must(template.ParseFiles("../web/template/manage_privileges.html"))
	t.Execute(w, data)
}
