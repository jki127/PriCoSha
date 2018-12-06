package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

// PageData is used for sending data to the template pages
type PageData struct {
	LoggedIn    bool
	Username    string
	Item        *b.ContentItem
	TaggedNames []*string
	Ratings     []*b.Rating
}

// getUserSession takes in a http.Request, reads the username cookie and
// returns two values:
// - a bool representing if the current user is logged in
// - a string representing the current user's username
//
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

func contentItemHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)
	itemID, err := strconv.Atoi(r.PostFormValue("itemID"))

	if err != nil {
		log.Println(err)
	}

	if !b.UserHasAccessToItem(username, itemID) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	pageData := PageData{
		LoggedIn:    logged,
		Username:    username,
		Item:        b.GetContentItemById(itemID),
		TaggedNames: b.GetTaggedByItemId(itemID),
		Ratings:     b.GetRatingsByItemId(itemID),
	}

	t := template.Must(template.ParseFiles("../web/template/content_item.html"))
	t.Execute(w, pageData)
}
