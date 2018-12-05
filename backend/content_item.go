package backend

import (
	"log"
	"time"
)

type ContentItem struct {
	ItemID   int
	Email    string
	FilePath string
	FileName string
	PostTime time.Time
	Fname    string
	Lname    string
}

type Rating struct {
	Email     string
	Rate_time time.Time
	Emoji     string
}

/*
GetPubContent queries DB for all Content_Item entities with a public
status and returns them as an array of ContentItem pointers.
*/
func GetPubContent() []*ContentItem {
	// Query DB for data
	rows, err := db.Query(`
	SELECT * FROM Content_Item
	WHERE is_pub = true
	AND post_time >= DATE_SUB(NOW(), INTERVAL 24 HOUR)
	`)
	if err != nil {
		log.Println(`backend: GetPubContent(): Could not
		query public content from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var (
		isPub int
		data  []*ContentItem
	)
	for rows.Next() {
		var CurrentItem ContentItem
		err = rows.Scan(&CurrentItem.ItemID, &CurrentItem.Email,
			&CurrentItem.FilePath, &CurrentItem.FileName,
			&CurrentItem.PostTime, &isPub)
		if err != nil {
			log.Println(`backend: GetPubContent(): Could not scan row data
			from public content query.`)
		}
		data = append(data, &CurrentItem)
	}

	return data
}

// GetUserContent retrieves all ContentItems from the DB that are
// - shared with the current user
// - posted by the current user
// - public content that has been posted in the past 24hrs
//
// The items are ordered in reverse chronological order
func GetUserContent(email string) []*ContentItem {
	rows, err := db.Query(`
	SELECT * FROM Content_Item
	WHERE item_id IN (
		-- All item ids shared in a user's friendgroups
		SELECT item_id FROM Share
		WHERE fg_name IN (
			-- All friend groups the user belongs to
			SELECT fg_name FROM Belong
			WHERE member_email=?
		)
	)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
	OR poster_email=?
	ORDER BY Content_Item.post_time DESC
	`, email, email)
	if err != nil {
		log.Println(`backend: GetPubContent(): Could not
		query public content from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var (
		isPub int
		data  []*ContentItem
	)
	for rows.Next() {
		var CurrentItem ContentItem
		err = rows.Scan(&CurrentItem.ItemID, &CurrentItem.Email,
			&CurrentItem.FilePath, &CurrentItem.FileName,
			&CurrentItem.PostTime, &isPub)
		if err != nil {
			log.Println(`backend: GetPubContent(): Could not scan row data
			from public content query.`)
		}
		data = append(data, &CurrentItem)
	}

	return data
}

func GetContentItemById(itemId int) *ContentItem {
	row := db.QueryRow(`
	SELECT item_id, poster_email, file_path, file_name, post_time, f_name,
	l_name FROM Content_Item
	JOIN Person ON Content_Item.poster_email = Person.email
	WHERE item_id = ?
	`, itemId)

	var item ContentItem
	err := row.Scan(&item.ItemID, &item.Email, &item.FilePath, &item.FileName,
		&item.PostTime, &item.Fname, &item.Lname)

	if err != nil {
		log.Println("GetContentItemById() scan error:", err)
	}

	return &item
}

func GetTaggedByItemId(itemId int) []*string {
	// This query concatenates first name and last name into one column `name`
	rows, err := db.Query(`
	SELECT CONCAT(f_name, " ", l_name) AS name FROM Tag
	JOIN Person ON Tag.tagged_email = Person.email
	WHERE item_id =? AND status = 1
	`, itemId)

	defer rows.Close()
	if err != nil {
		log.Println("GetTaggedByItemId() query error: ", err)
	}

	var (
		name        string
		taggedNames []*string
	)
	for rows.Next() {
		rows.Scan(&name)
		taggedNames = append(taggedNames, &name)
	}

	return taggedNames
}

func GetRatingsByItemId(itemId int) []*Rating {
	rows, err := db.Query(`
	SELECT email, rate_time, emoji FROM Rate WHERE item_id=?
	`, itemId)

	defer rows.Close()
	if err != nil {
		log.Println("GetRatingsByItemId() query error: ", err)
	}

	var (
		rating  Rating
		ratings []*Rating
	)
	for rows.Next() {
		rows.Scan(&rating.Email, &rating.Rate_time, &rating.Emoji)
		ratings = append(ratings, &rating)
	}

	return ratings
}