package main

import (
	"html/template"

	"net/http"
	b "pricosha/backend"
)
func duplicateAddFriendHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "addFriendErr")

	url := r.URL
	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
	fname := "example"
	lname := "example"

	data := struct {
		FGName     string
		OwnerEmail string
		UserEmails	[]*string
	}{
		fgName,
		ownerEmail,
		b.GetEmail(fname, lname),
	}



	//t := template.Must(template.New("").ParseFiles("../web/template/friend_groups.html",
	//	"../web/template/base.html"))
	//t.ExecuteTemplate(w, "base", CurrentFGD)
}