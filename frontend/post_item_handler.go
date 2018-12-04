package main

import (
	"log"
	_ "log"
	"net/http"
	b "pricosha/backend"
	"strings"
	"time"
)

// Handles requests to post content item data
func postItemHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("username")
	var username string
	if err != nil {
		// User is not logged in, redirect
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	username = cookie.Value

	r.ParseForm()
	for key, value := range r.Form {
		log.Println(key, value)
	}

	NewContentItem := b.ContentItem{
		ItemID:   b.GetNewItemID(),
		Email:    username,
		FilePath: r.FormValue("filePath"),
		FileName: r.FormValue("itemName"),
		PostTime: time.Now(),
	}

	var isPub int
	privacySetting := r.FormValue("shareSetting")
	if privacySetting == "public" {
		isPub = 1
	} else {
		isPub = 0
	}
	b.ExecInsertContentItem(NewContentItem, isPub)

	// If the content item is private, need to update Share table for each FriendGroup
	if isPub == 0 {
		sharedGroups := strings.Fields(r.FormValue("friendGroup"))
		for group := range sharedGroups {
			groupInfo := strings.Split(sharedGroups[group], "_")
			SharedGroup := b.FriendGroup{
				MemberEmail: username,
				FGName:      groupInfo[0],
				OwnerEmail:  groupInfo[1],
			}
			b.ExecInsertSharedContentItemToGroup(SharedGroup.FGName, SharedGroup.OwnerEmail, NewContentItem.ItemID)
			// right now only executes/processes the first FriendGroup returned by the form
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return
}
