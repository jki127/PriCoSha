package backend

import (
	"log"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

// Tag holds info of Tag items, taken from database data
type Tag struct {
	TaggerEmail string
	TaggedEmail string
	ItemID      int
	TagTime     string
	FileName    string
	FilePath    string
	TagID       int
}

//GetPendingTags Gathers the Pending tags a user has by searching for tags with status=false, where the user is the Tagged person
func GetPendingTags(username string) []*Tag {
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path 
		FROM Tag NATURAL JOIN Content_Item
		WHERE status = false
		AND tagged_email=?`,
		username)
	if err != nil {
		log.Println(`backend: GetPendingTags(): Could not
		query tags content from DB.`)
	}
	defer rows.Close()

	var data []*Tag
	tagCounter := 0
	for rows.Next() {
		var CurrentTag Tag
		err = rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
			&CurrentTag.ItemID, &CurrentTag.TagTime,
			&CurrentTag.FileName, &CurrentTag.FilePath)
		if err != nil {
			log.Println(`backend: GetPendingTags(): Could not scan row data
			from tag content query.`)
		}
		CurrentTag.TagID = tagCounter
		data = append(data, &CurrentTag)
		tagCounter++
	}

	return data
}

//GetAcceptedTags Gathers the Pending tags a user has by searching for tags with status=true, where the user is the Tagged person
func GetAcceptedTags(username string) []*Tag {
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path 
		FROM Tag NATURAL JOIN Content_Item
		WHERE status = true
		AND tagged_email=?`,
		username)
	if err != nil {
		log.Println(`backend: AcceptedTags(): Could not
		query tags content from DB.`)
	}
	defer rows.Close()

	var data []*Tag

	for rows.Next() {
		var CurrentTag Tag
		err = rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
			&CurrentTag.ItemID, &CurrentTag.TagTime,
			&CurrentTag.FileName, &CurrentTag.FilePath)
		if err != nil {
			log.Println(`backend: GetAcceptedTags(): Could not scan row data
			from tag content query.`)
		}
		data = append(data, &CurrentTag)
	}

	return data
}

//DeclineTag Deletes tag with given parameters from the database
func DeclineTag(tagger string, tagged string, itemID int) {
	_, err := db.Exec(`DELETE FROM Tag 
	WHERE tagger_email=? 
	AND tagged_email=? 
	AND item_id=?`,
		tagger, tagged, itemID)
	if err != nil {
		log.Println(`backend: DeclineTag(): Could not
		Delete from DB.`, tagger, tagged, itemID)
		return
	}
	log.Println("Tag Declined succesfully")
}

//AcceptTag Updates tag with given paramaters as status=true
func AcceptTag(tagger string, tagged string, itemID int) {
	_, err := db.Exec(`UPDATE Tag SET status=1 
	WHERE tagger_email=? 
	AND tagged_email=? 
	AND item_id=?`,
		tagger, tagged, itemID)

	if err != nil {
		log.Println(`backend: AcceptTag(): Could not
		UPDATE DB.`, tagger, tagged, itemID)
		return
	}
	log.Println("Tag Accepted succesfully")
}
