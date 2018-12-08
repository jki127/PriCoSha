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
