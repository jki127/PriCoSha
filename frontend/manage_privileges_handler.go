package main

import (
	"html/template"
	"log"
	"net/http"
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

	if username != owner {
		log.Println(`User is not the owner of the group and cannot manage
			privileges.`)
		http.Redirect(w, r, "/friendgroups", http.StatusFound)
		return
	}

	group := queryData["fgn"][0]

	data := struct {
		Logged     bool
		FGName     string
		OwnerEmail string
	}{
		true,
		group,
		owner,
	}

	t := template.Must(template.ParseFiles("../web/template/manage_privileges.html"))
	t.Execute(w, data)
}
