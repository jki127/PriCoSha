package backend

import (
	"log"
	"time"
)

func AddRatingToDB(emoji string, username string, itemID int) {
	_, err := db.Exec(`INSERT into Rate VALUES (?,?,?,?)`, username, itemID, time.Now(), emoji)
	if err != nil {
		log.Println(`backend: addRatingToDB(): Could not
		INSERT rating into DB.`)
		return
	}
	log.Println("Rating added successfully")
}
