package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
)

// Handles requests to validate user data and sets cookies accordingly
func validateLoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err == nil {
		// User is already logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	// This check should be revised later
	if username == "" || password == "" {
		cookie := http.Cookie{Name: "logErr", Value: "empty"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if ok := b.ValidateInfo(username, password); ok {
		log.Println("User logged in with:", username, password)
		// Set cookie with user info
		cookie := http.Cookie{Name: "username", Value: username}
		http.SetCookie(w, &cookie)
		// Delete logErr cookie if it exists
		clearCookie(&w, r, "logErr")

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	log.Println("User failed to log in with:", username, password)
	cookie := http.Cookie{Name: "logErr", Value: "fail"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
	return
}
