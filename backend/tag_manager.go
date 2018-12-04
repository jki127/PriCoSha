package backend

import (
	"database/sql"
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
}

//getTagRows Queries the database to find rows in which the user is tagged and the status matches the given parameter
func getTagRows(status bool, username string) *sql.Rows {
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path 
		FROM Tag NATURAL JOIN Content_Item
		WHERE status=?
		AND tagged_email=?`,
		status, username)
	if err != nil {
		log.Println(`backend: getTagRows(): Could not
		query tags content from DB.`)
	}
	return rows
}

//GetPendingTags Gathers the Pending tags a user has by searching for tags with status=false, where the user is the Tagged person
func GetPendingTags(username string) []*Tag {
	rows := getTagRows(false, username)
	defer rows.Close()

	var data []*Tag
	for rows.Next() {
		var CurrentTag Tag
		err := rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
			&CurrentTag.ItemID, &CurrentTag.TagTime,
			&CurrentTag.FileName, &CurrentTag.FilePath)
		if err != nil {
			log.Println(`backend: GetPendingTags(): Could not scan row data
			from tag content query.`)
		}
		data = append(data, &CurrentTag)
	}

	return data
}

//GetAcceptedTags Gathers the Pending tags a user has by searching for tags with status=true, where the user is the Tagged person
func GetAcceptedTags(username string) []*Tag {
	rows := getTagRows(true, username)
	defer rows.Close()

	var data []*Tag

	for rows.Next() {
		var CurrentTag Tag
		err := rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
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
