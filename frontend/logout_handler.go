package main

import (
	"log"
	"net/http"
	"time"
)

// Handles requests to logout and deletes cookies accordingly
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		log.Println("User was not logged in and cannot be logged out.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	cookie := http.Cookie{
		Name:    "username",
		Value:   "",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
