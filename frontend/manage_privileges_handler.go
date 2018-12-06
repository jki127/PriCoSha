package main

import (
	"html/template"
	"log"
	"net/http"

	b "pricosha/backend"
)

func managePrivilegesHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "chngPrivErr")

	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println("User was not logged in and cannot manage privileges.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := cookie.Value

	url := r.URL
	queryData := url.Query()
	owner := queryData["oe"][0]
	group := queryData["fgn"][0]
	role := b.GetRole(group, owner, username)

	if role == 2 {
		log.Println(`User is only a member of the group and cannot manage
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
