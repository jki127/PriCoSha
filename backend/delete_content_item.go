package backend

import (
	"log"
)

/*
ExecDeleteContentItem prepares and executes statement to delete content item from database
*/
func ExecDeleteContentItem(ID int64) {
	
	// delete content item from Content_Item
	statement, err := db.Prepare(`DELETE FROM Content_Item WHERE item_id = ?`)
	if err != nil {
		log.Println(`delete_content_item: execDeleteContentItem(): Could not prepare 
			content item deletion`)
		log.Println(err)
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		log.Println(`delete_content_item: execDeleteContentItem(): Could not execute 
			content item deletion`)
		log.Println(err)
	}
}