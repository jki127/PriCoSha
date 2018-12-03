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

	// Establish functions for handling requests to specific pages
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateLoginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/tag_manager",tagManagerHandler)
	http.HandleFunc("/decline",declineTagHandler)
	http.HandleFunc("/accept",acceptTagHandler)



	// Start server
	log.Println("Frontend spun up!")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("frontend: main(): ListenAndServe() failed.")
	}
}
