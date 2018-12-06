package main

import (
	"net/http"
	"time"
)

// If the user is not logged in then it will return the values (false, "")
func getUserSessionInfo(r *http.Request) (bool, string) {
	cookie, err := r.Cookie("username")
	var logged bool
	var username string
	if err != nil {
		logged = false
		username = ""
	} else {
		logged = true
		username = cookie.Value
	}
	return logged, username
}

// Will attempt to clear any cookie at given string value
func clearCookie(w *http.ResponseWriter, r *http.Request, name string) {
	if _, err := r.Cookie(name); err == nil {
		c := http.Cookie{
			Name:    name,
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(*w, &c)
	}
}
