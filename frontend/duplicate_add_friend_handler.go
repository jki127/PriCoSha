package main

import (
	"html/template"

	"net/http"
	b "pricosha/backend"
)

type DFD struct {
	FGName     string
	OwnerEmail string
	UserEmails	[]*string
}
func duplicateAddFriendHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "addFriendErr")

	url := r.URL
	queryData := url.Query()
	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
	fname := queryData["fname"][0]
	lname := queryData["lname"][0]

	CurrentDFD := DFD{
		FGName:			fgName,
		OwnerEmail:		ownerEmail,
		UserEmails:		b.GetEmail(fname, lname),
	}


	t := template.Must(template.ParseFiles("../web/template/duplicate_friend.html"))
	t.Execute(w, CurrentDFD)
}