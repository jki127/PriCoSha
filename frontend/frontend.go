package main

import (
	"log"
	"net/http"

	b "pricosha/backend"
)

// Port that server listens to http requests on (only edit number value)
var httpPort = ":" + "8080"

func main() {
	if b.TestDB() == nil {
		log.Println("Database connected successfully!")
	} else {
		log.Fatal("Database could not be contacted.")
	}

	// Serve CSS files from assets/css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../assets/css/"))))

	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	http.HandleFunc("/item", contentItemHandler)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateLoginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/postItem", postItemHandler)

	http.HandleFunc("/addtag", addTagHandler)

	http.HandleFunc("/friendgroups", friendGroupHandler)
	http.HandleFunc("/formAddFriend", formAddFriendHandler)
	http.HandleFunc("/addFriend", addFriendHandler)

	http.HandleFunc("/deleteFriend", deleteFriendHandler)
	http.HandleFunc("/formDeleteFriend", formDeleteFriendHandler)

	http.HandleFunc("/tag_manager", manageTagHandler)
	http.HandleFunc("/decline", declineTagHandler)
	http.HandleFunc("/accept", acceptTagHandler)

	http.HandleFunc("/profile", profileHandler)

	http.HandleFunc("/managePrivilege", managePrivilegesHandler)
	http.HandleFunc("/changePrivilege", changePrivilegeHandler)
	http.HandleFunc("/unshare", unshareHandler)
	http.HandleFunc("/renameGroup", renameGroupHandler)
	http.HandleFunc("/changeOwner", changeOwnerHandler)
	http.HandleFunc("/deleteGroup", deleteGroupHandler)

  http.HandleFunc("/duplicateNames", duplicateAddFriendHandler)
	http.HandleFunc("/chooseName", chooseAddedFriendHandler)

	// Start server
	log.Println("Frontend spun up!")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("frontend: main(): ListenAndServe() failed.")
	}
}
