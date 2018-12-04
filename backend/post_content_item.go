package backend

import (
	"database/sql"
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
		log.Println(`post_content_item: GetFriendUserGroup(username string): Could not
		query user's Friend Groups from DB.`)
	}
	defer rows.Close()

	var data []*FriendGroup

	for rows.Next() {
		var CurrentGroup FriendGroup
		err = rows.Scan(&CurrentGroup.MemberEmail, &CurrentGroup.FGName,
			&CurrentGroup.OwnerEmail)
		if err != nil {
			log.Println(`post_content_item: GetFriendUserGroup(username string): Could not scan row data
			from public content query.`)
		}
		data = append(data, &CurrentGroup)
	}
	return data
}

func GetNewItemID() int {
	// Query DB for current max item ID
	row := db.QueryRow(`SELECT MAX(item_id) FROM Content_Item`)

	var maxItemID int
	err := row.Scan(&maxItemID)

	switch {
	case err == sql.ErrNoRows:
		log.Println("post_content_item: GetNewItemID(): no item IDs found")
		return -1
	case err != nil:
		log.Println("post_content_item: GetNewItemID(): non nil Scan() error")
		return -1
	default:
		return maxItemID + 1
	}
}

func ExecInsertContentItem(item ContentItem, isPub int) {
	log.Println("post_content_item: id:", item.ItemID, "privacy setting:", isPub)

	statement, err := db.Prepare(`INSERT INTO Content_Item VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println("post_content_item: execInsertContentItem(): Could not prepare content item insertion")
	}
	defer statement.Close()

	_, err = statement.Exec(item.ItemID, item.Email, item.FilePath, item.FileName, item.PostTime, isPub)
	if err != nil {
		log.Println("post_content_item: execInsertContentItem(): Could not execute content item insertion")
		log.Println(err)
	}
}

func ExecInsertSharedContentItemToGroup(FGName string, OwnerEmail string, itemID int) {
	statement, err := db.Prepare(`INSERT INTO Share VALUES (?, ?, ?)`)
	if err != nil {
		log.Println("post_content_item: ExecInsertSharedContentItemToGroup(): Could not prepare shared content insertion")
	}
	defer statement.Close()

	_, err = statement.Exec(FGName, OwnerEmail, itemID)
	if err != nil {
		log.Println("post_content_item: ExecInsertSharedContentItemToGroup(): Could not execute shared content insertion")
		log.Println(err)
	}
}
