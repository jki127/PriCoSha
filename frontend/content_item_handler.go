package main

import (
	"html/template"
	"log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

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

	// Gets groups the user can remove this item from
	removes := b.UserCanRemoveFrom(username, itemID)

	pageData := struct {
		Logged      bool
		Username    string
		Item        *b.ContentItem
		TaggedNames []*string
		Ratings     []*b.Rating
		Comments    []*b.Comment
		Removes     []*b.FriendGroup
	}{
		logged,
		username,
		b.GetContentItemById(itemID),
		b.GetTaggedByItemId(itemID),
		b.GetRatingsByItemId(itemID),
		b.GetCommentsByItemId(itemID),
		removes,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/content_item.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", pageData)
}
