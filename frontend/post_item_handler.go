package main

import (
	_ "log"
	"net/http"
	_ "pricosha/backend"
	_ "time"
)

// Handles requests to post content item data
func postItemHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
