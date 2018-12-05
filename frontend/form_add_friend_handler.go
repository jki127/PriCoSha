package main

import (
	"html/template"
	"net/http"
)

func formAddFriendHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	queryData := url.Query()

	cookie, err := r.Cookie("addFriendErr")
	var errMsg string
	var isErr bool
	if err == nil {
		isErr = true
		if cookie.Value == "empty" {
			errMsg = "Data fields were left empty. Please retry."
		} else if cookie.Value == "duplicates" {
			errMsg = "Multiple people have the same name. Please specify."
		} else if cookie.Value == "nonexistent" {
			errMsg = "Person you are trying to add does not exist. Please retry."
		} else if cookie.Value == "alreadyBelongs" {
			errMsg = "Person you are trying to add already belongs to Friend Group. Please retry."
		}
	} else {
		isErr = false
	}

	fgName := queryData["fgn"][0]
	ownerEmail := queryData["oe"][0]
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

	t := template.Must(template.ParseFiles("../web/template/add_friend.html"))
	t.Execute(w, data)
}
