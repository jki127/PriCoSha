package backend

import "log"

// GetFolders gets all the folders that a user has created
func GetFolders(email string) []*Folder {
	rows, err := db.Query(`SELECT folder_name FROM Folder WHERE email =?`, email)
	if err != nil {
		log.Println("content_folder: GetFolders(): Could not query folder")
	}
	defer rows.Close()

	var folders []*Folder

	for rows.Next() {
		var folder Folder
		rows.Scan(&folder.Name)
		folders = append(folders, &folder)
	}

	return folders
}

// GetContentInFolder gets all the content items with a folder based on
// the specified folder name and user email
func GetContentInFolder(folderName string, email string) []*ContentItem {
	rows, err := db.Query(`
	SELECT item_id, poster_email, file_path, file_name, post_time FROM Include
	NATURAL JOIN Content_Item
	WHERE folder_name = ? AND email = ?
	`, folderName, email)

	if err != nil {
		log.Println(`content_folder: GetContentInFolder(): Could not query db for
			content items in folder`)
	}

	var items []*ContentItem

	for rows.Next() {
		var item ContentItem
		rows.Scan(&item.ItemID, &item.Email, &item.FilePath, &item.FileName,
			&item.PostTime)
		items = append(items, &item)
	}

	return items
}

func CreateFolder(folderName string, email string) error {
	statement, err := db.Prepare(`
	INSERT INTO Folder (folder_name, email)
	VALUES (?, ?)
	`)
	if err != nil {
		log.Println("content_folder: createFolder(): could not prepare insert folder:", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(folderName, email)
	if err != nil {
		log.Println("content_folder: createFolder(): could not insert folder", err)
		return err
	}

	return nil
}

// GetContentNotInFolder gets all the content items a user has access to that
// are currently not in the current folder
func GetContentNotInFolder(folderName string, email string) []*ContentItem {
	rows, err := db.Query(`
	SELECT item_id, poster_email, file_path, file_name
	FROM Content_Item
	WHERE (item_id IN (
		-- All item ids shared in a user's friendgroups
		SELECT item_id FROM Share
		WHERE (fg_name, owner_email) IN (
			-- All friend groups the user belongs to
			SELECT fg_name, owner_email FROM Belong
			WHERE member_email=?
		)
	)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
	OR poster_email=?) AND
	item_id NOT IN (
		SELECT item_id FROM Include
		NATURAL JOIN Content_Item
		WHERE folder_name = ? AND email = ?
	)
	ORDER BY Content_Item.post_time DESC
	`, email, email, folderName, email)

	if err != nil {
		log.Println(`content_folder: GetContentInFolder(): Could not query db for
			content items in folder`)
	}

	var items []*ContentItem

	for rows.Next() {
		var CurrentItem ContentItem
		err := rows.Scan(&CurrentItem.ItemID, &CurrentItem.Email,
			&CurrentItem.FilePath, &CurrentItem.FileName)

		if err != nil {
			log.Println("content_folder: GetContentNotInFolder(): Could not scan rows")
		}

		items = append(items, &CurrentItem)
	}

	return items
}

func AddItemToFolder(folderName string, email string, itemID int) {
	statement, err := db.Prepare(`
	INSERT INTO Include (folder_name, email, item_id)
	VALUES (?, ?, ?)
	`)
	if err != nil {
		log.Println("content_folder: AddItemToFolder(): could not prepare insert Include:", err)
	}
	defer statement.Close()

	_, err = statement.Exec(folderName, email, itemID)
	if err != nil {
		log.Println("content_folder: AddItemToFolder(): could not insert Include", err)
	}
}
