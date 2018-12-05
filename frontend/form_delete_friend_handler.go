package main

import (
	"html/template"
	"net/http"
)

func formDeleteFriendHandler(w http.ResponseWriter, r *http.Request){
	url := r.URL
	queryData := url.Query()

	cookie, err := r.Cookie("deleteFriendErr")
	var errMsg string
	var isErr bool

	if err == nil {
		isErr = true
		if cookie.Value == "empty" {
			errMsg = "Data fields were left empty. Please retry."
		} else if cookie.Value == "nonexistent" {
			errMsg = "Person you are trying to add does not exist. Please retry."
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

	t := template.Must(template.ParseFiles("../web/template/delete_friend.html"))
	t.Execute(w, data)
}