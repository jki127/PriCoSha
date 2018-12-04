package backend

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// TagItem holds info of Tag entities in the database
type TagItem struct {
	TaggerEmail string
	TaggedEmail string
	ItemID      int
	Status      bool // false = private; true = public
	TagTime     time.Time
}

// execInsertTag takes the data for a tag entity and inserts in into the table
func execInsertTag(id int, uTagger string, uTagged string, pubVal bool) {
	log.Println("backend: id:", id, "uTagger:", uTagger,
		"uTagged:", uTagged, "pubVal:", pubVal)

	statement, err := db.Prepare(`INSERT Tag VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println("backend: InsertTag(): Could not prepare tag insertion")
	}
	defer statement.Close()

	_, err = statement.Exec(uTagger, uTagged, id, pubVal, time.Now())
	if err != nil {
		log.Println("backend: InsertTag(): Could not execute tag insertion")
		log.Println(err)
	}
}

/*
InsertTag receives tag info for insertion to the table, checks the validity of the
insertion, and then calls execInsertTag with valid info if necessary. Otherwise, it
returns an error for the frontend if the insertion could not be completed.
*/
func InsertTag(id int, uTagger string, uTagged string) error {
	log.Println("backend: id:", id, "uTagger:", uTagger, "uTagged:", uTagged)
	if uTagger == uTagged {
		execInsertTag(id, uTagger, uTagged, true)
		return nil
	}

	/*
		This query is meant more to just return a row than to return any specific
		content. If it returns, that means the following: that an item exists with
		item_id=id; that that item exists within the set of public items OR that the
		member who was tagged is a member of one of the friend groups that item is
		shared in.
		If the query returns a row, that means the user can view the content item
		they were tagged publically or privately.
	*/
	row := db.QueryRow(`SELECT item_id
		FROM Content_Item LEFT OUTER JOIN Share
		USING (item_id)
		WHERE item_id=?
		AND (
			item_id IN (
				SELECT item_id
				FROM Content_Item
				WHERE is_pub=true
			)
			OR (fg_name, owner_email) IN (
				SELECT fg_name, owner_email
				FROM Belong
				WHERE member_email=?
			)
		)`, id, uTagged)

	var itemID int

	err := row.Scan(&itemID)

	fmt.Println(itemID)

	switch {
	case err == sql.ErrNoRows:
		returnErr := errors.New("noview")
		return returnErr
	case err != nil:
		returnErr := errors.New("failed")
		return returnErr
	default:
		execInsertTag(id, uTagger, uTagged, false)
		return nil
	}
}
