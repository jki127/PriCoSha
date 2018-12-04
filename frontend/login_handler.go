package main

import (
	"html/template"
	"net/http"
)

// Handles requests to login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err == nil {
		// User is already logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Check if there was a previous login error
	cookie, err := r.Cookie("logErr")
	var errMsg string
	var isErr bool
	if err == nil {
		isErr = true
		if cookie.Value == "empty" {
			errMsg = "Data fields were left empty. Please retry."
		} else if cookie.Value == "fail" {
			errMsg = "Incorrect login data. Please retry."
		}
	} else {
		isErr = false
	}

	data := struct {
		IsErr  bool
		ErrMsg string
	}{
		isErr,
		errMsg,
	}

	t := template.Must(template.ParseFiles("../web/template/login.html"))
	t.Execute(w, data)
}

