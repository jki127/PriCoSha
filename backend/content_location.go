package backend

import (
	"log"
)

// GetUserLocations takes in the current user's email address and returns
// a map of locations the user has been to
// The keys are the name of the location and the values are the number of
// content items from that location
func GetUserLocations(email string) map[string]int {
	rows, err := db.Query(`
	SELECT location, count(item_id) FROM Content_Item
	WHERE (item_id IN (
		-- All item ids shared in a user's friendgroups
		SELECT item_id FROM Share
		WHERE (fg_name, owner_email) IN (
			-- All friend groups the user belongs to
			SELECT fg_name, owner_email FROM Belong
			WHERE member_email= ?
		)
	)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
	OR poster_email = ?)
	AND location IS NOT NULL
	GROUP BY location
	`, email, email)

	if err != nil {
		log.Println("getUserLocations(): Could not query Content_Item")
	}
	defer rows.Close()

	locations := make(map[string]int)

	for rows.Next() {
		var location string
		var count int
		rows.Scan(&location, &count)
		locations[location] = count
	}

	return locations
}

// GetUserContentByLocation takes in the current user's email and a location
// and returns a list of content items that are from that location and that the
// user has access to
func GetUserContentByLocation(email string, location string) []*ContentItem {
	rows, err := db.Query(`
	SELECT item_id, poster_email, file_path, file_name, post_time
	FROM Content_Item
	WHERE location = ? AND (item_id IN (
		-- All item ids shared in a user's friendgroups
		SELECT item_id FROM Share
		WHERE (fg_name, owner_email) IN (
			-- All friend groups the user belongs to
			SELECT fg_name, owner_email FROM Belong
			WHERE member_email = ?
		)
	)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
	OR poster_email = ?)
	ORDER BY Content_Item.post_time DESC
	`, location, email, email)
	if err != nil {
		log.Println(`backend: GetPubContent(): Could not
		query public content from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var (
		items []*ContentItem
	)
	for rows.Next() {
		var currentItem ContentItem
		err = rows.Scan(&currentItem.ItemID, &currentItem.Email,
			&currentItem.FilePath, &currentItem.FileName,
			&currentItem.PostTime)
		if err != nil {
			log.Println(`backend: GetPubContent(): Could not scan row data
			from public content query.`)
		}
		items = append(items, &currentItem)
	}

	return items
}
