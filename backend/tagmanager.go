package backend

import (
	"log"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

// PendingTag holds info of Pending Tag items, taken from database data
type PendingTag struct {
	TaggerEmail string
	TaggedEmail string
	ItemID int
	TagTime string
	FileName string
	FilePath string
	TagID int
}
// Accepted holds info of Pending Tag items, taken from database data
type AcceptedTag struct {
	TaggerEmail string
	TaggedEmail string
	ItemID int
	TagTime string
	FileName string
	FilePath string
	TagID int
}

func GetPendingTags(username string) []*PendingTag{
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path FROM Tag NATURAL JOIN Content_Item
		WHERE status = false
		AND tagged_email=?`,username)
	if err != nil {
		log.Println(`backend: GetPendingTags(): Could not
		query tags content from DB.`)
	}
	defer rows.Close()
	
	var data []*PendingTag
	tagCounter:=0
	for rows.Next() {
		var CurrentTag PendingTag
		err = rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
			&CurrentTag.ItemID, &CurrentTag.TagTime,
			&CurrentTag.FileName, &CurrentTag.FilePath)
		if err != nil {
			log.Println(`backend: GetPendingTags(): Could not scan row data
			from tag content query.`)
		}
		CurrentTag.TagID=tagCounter
		data = append(data, &CurrentTag)
		tagCounter+=1
	}

	return data
}

func GetAcceptedTags(username string) []*AcceptedTag{
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path FROM Tag NATURAL JOIN Content_Item
		WHERE status = true
		AND tagged_email=?`,username)
	if err != nil {
		log.Println(`backend: AcceptedTags(): Could not
		query tags content from DB.`)
	}
	defer rows.Close()
	
	var data []*AcceptedTag

	for rows.Next() {
		var CurrentTag AcceptedTag
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
func  DeclineTag(tagger string, tagged string, itemID int){
	_, err := db.Exec(`DELETE FROM Tag WHERE tagger_email=? AND tagged_email=? AND item_id=?`,tagger,tagged,itemID)
	if err != nil {
		log.Println(`backend: DeclineTag(): Could not
		Delete from DB.`,tagger,tagged,itemID)
		return
	}
		println("Tag Declined succesfully")
	
	return
}	
func  AcceptTag(tagger string, tagged string, itemID int){
	_, err := db.Exec(`UPDATE Tag SET status=1 WHERE tagger_email=? AND tagged_email=? AND item_id=?`,tagger,tagged,itemID)
	if err != nil {
		log.Println(`backend: AcceptTag(): Could not
		UPDATE DB.`,tagger,tagged,itemID)
		return
	}
		println("Tag Accepted succesfully")
	
	return
}	
