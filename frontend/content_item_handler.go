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

func contentItemHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)
	urlParams := r.URL.Query()
	itemID, err := strconv.Atoi(urlParams["iid"][0])
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
