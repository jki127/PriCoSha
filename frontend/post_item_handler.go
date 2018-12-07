package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	b "pricosha/backend"
)

// Handles requests to post content item data
func postItemHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		// User is not logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	username := cookie.Value

	// Parse HTML form for user-entered data about content item
	r.ParseForm()
	log.Println("New Content Item info:")
	for key, value := range r.Form {
		log.Println(key, value)
	}

	// Create new ContentItem with appropriate data
	NewContentItem := b.ContentItem{
		Email:    username,
		FilePath: r.FormValue("filePath"),
		FileName: r.FormValue("itemName"),
		PostTime: time.Now(),
	}

	// Designate privacy setting
	var isPub int
	privacySetting := r.FormValue("shareSetting")
	if privacySetting == "public" {
		isPub = 1
	} else {
		isPub = 0
	}
	// Send info to backend to be inserted into database
	lastID := b.ExecInsertContentItem(NewContentItem, isPub)
	log.Println("Inserted a Content_Item into the db!")

	// If the content item is private, need to update Share table for each FriendGroup
	if isPub == 0 {
		sharedGroups := r.Form["friendGroup"]
		// Create a FriendGroup for each chosen group to share item with
		for group := range sharedGroups {
			groupInfo := strings.Split(sharedGroups[group], "_")
			SharedGroup := b.FriendGroup{
				MemberEmail: username,
				FGName:      groupInfo[0],
				OwnerEmail:  groupInfo[1],
			}
			// Send info to backend to be inserted into database
			b.ExecInsertSharedContentItemToGroup(SharedGroup.FGName,
				SharedGroup.OwnerEmail, lastID)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
