package main

import (
	"html/template"
	"net/http"
)

// Handles requests to upload content item page
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// Check and make sure user is logged in
	_, err := r.Cookie("username")
	if err != nil {
		// User is not logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	t := template.Must(template.ParseFiles("../web/template/upload.html"))
	t.Execute(w, nil)
}