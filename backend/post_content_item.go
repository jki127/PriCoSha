package backend

import (
	"log"
)

// FriendGroup holds info of Friend_Group entities in the database
type FriendGroup struct {
	MemberEmail string
	FGName      string
	OwnerEmail  string
}

/*
GetUserFriendGroup receives info from database about FriendGroups that user belongs to
*/
func GetUserFriendGroup(username string) []*FriendGroup {
	// Query DB for data
	rows, err := db.Query(`SELECT * FROM Belong WHERE member_email=?`, username)
	if err != nil {
		log.Println(`post_content_item: GetFriendUserGroup(): Could not query user's 
		Friend Groups from DB.`)
	}
	defer rows.Close()

	var data []*FriendGroup

	for rows.Next() {
		var CurrentGroup FriendGroup
		err = rows.Scan(&CurrentGroup.MemberEmail, &CurrentGroup.FGName,
			&CurrentGroup.OwnerEmail)
		if err != nil {
			log.Println(`post_content_item: GetFriendUserGroup(): Could not scan row 
			data from public content query.`)
		}
		data = append(data, &CurrentGroup)
	}
	return data
}

/*
ExecInsertContentItem prepares and executes statement to insert new item into
Content_Item
*/
func ExecInsertContentItem(item ContentItem, isPub int) int64 {
	statement, err := db.Prepare(`INSERT INTO Content_Item (poster_email, file_path,
		file_name, post_time, is_pub) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(`post_content_item: execInsertContentItem(): Could not prepare 
			content item insertion`)
		log.Println(err)
	}
	defer statement.Close()

	row, err := statement.Exec(item.Email, item.FilePath, item.FileName,
		item.PostTime, isPub)
	if err != nil {
		log.Println(`post_content_item: execInsertContentItem(): Could not execute 
			content item insertion`)
		log.Println(err)
	}

	id, err := row.LastInsertId()
	if err != nil {
		log.Println(`post_content_item: execInsertContentItem(): Could not read
			id of last insert`)
	}

	return id
}

/*
ExecInsertSharedContentItemToGroup prepares and executes statement to insert info
about privately shared content items to Share
*/
func ExecInsertSharedContentItemToGroup(FGName string, OwnerEmail string, itemID int64) {
	log.Println("post_content_item: id:", itemID, "privately shared with", FGName,
		"owned by", OwnerEmail)
	statement, err := db.Prepare(`INSERT INTO Share VALUES (?, ?, ?)`)
	if err != nil {
		log.Println(`post_content_item: ExecInsertSharedContentItemToGroup(): Could '
			not prepare shared content insertion`)
	}
	defer statement.Close()

	_, err = statement.Exec(FGName, OwnerEmail, itemID)
	if err != nil {
		log.Println(`post_content_item: ExecInsertSharedContentItemToGroup(): Could 
			not execute shared content insertion`)
		log.Println(err)
	}
}
