package backend

import (
	"log"
	"math/rand"
)

/*
GetPubContent queries DB for all Content_Item entities with a public
status and returns them as an array of ContentItem pointers.
*/
func GetPubContent() []*ContentItem {
	// Query DB for data
	rows, err := db.Query(`
	SELECT item_id, poster_email, file_path, file_name, post_time,
	is_pub FROM Content_Item
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
		CurrentItem.RandImg = rand.Intn(8)

		if CheckPoll(CurrentItem.ItemID) {
			CurrentItem.IsPoll = true
			CurrentItem.Votes = GetVotes(CurrentItem.ItemID)
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
	SELECT item_id, poster_email, file_path, file_name, post_time, is_pub, f_name, l_name 
	FROM Content_Item JOIN Person ON Content_Item.poster_email=Person.email
	WHERE item_id IN (
		-- All item ids shared in a user's friendgroups
		SELECT item_id FROM Share
		WHERE (fg_name, owner_email) IN (
			-- All friend groups the user belongs to
			SELECT fg_name, owner_email FROM Belong
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
			&CurrentItem.PostTime, &isPub, &CurrentItem.Fname, &CurrentItem.Lname)
		if err != nil {
			log.Println(`backend: GetPubContent(): Could not scan row data
			from public content query.`)
		}
		CurrentItem.RandImg = rand.Intn(8)
		CurrentItem.Comments = GetCommentsByItemId(CurrentItem.ItemID)
		CurrentItem.Ratings = GetRatingsByItemId(CurrentItem.ItemID)

		if CheckPoll(CurrentItem.ItemID) {
			CurrentItem.IsPoll = true
			CurrentItem.Votes = GetVotes(CurrentItem.ItemID)
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
	if CheckPoll(item.ItemID) {
		item.IsPoll = true
		item.Votes = GetVotes(item.ItemID)
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

	var taggedNames []*string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		taggedNames = append(taggedNames, &name)
	}

	return taggedNames
}

func GetCommentsByItemId(itemId int) []*Comment {
	rows, err := db.Query(`
	SELECT Comment.email, comment_time, body, f_name, l_name FROM Comment JOIN Person ON Comment.email=Person.email WHERE item_id=?
	ORDER BY comment_time DESC`, itemId)

	defer rows.Close()
	if err != nil {
		log.Println("content_item: GetCommentsByItemId() query error: ", err)
	}

	var comments []*Comment
	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.Email, &comment.CommentTime, &comment.Body, &comment.Fname, &comment.Lname)
		comments = append(comments, &comment)
	}

	return comments
}

func GetRatingsByItemId(itemId int) []*Rating {
	rows, err := db.Query(`
	SELECT email, rate_time, emoji FROM Rate WHERE item_id=?
	ORDER BY rate_time DESC`, itemId)

	defer rows.Close()
	if err != nil {
		log.Println("content_item: GetRatingsByItemId() query error: ", err)
	}

	var ratings []*Rating
	for rows.Next() {
		var rating Rating
		rows.Scan(&rating.Email, &rating.Rate_time, &rating.Emoji)
		ratings = append(ratings, &rating)
	}

	return ratings
}

// UserHasAccessToItem checks to see if the current user, specified by username,
// is any friend groups that have access the current content item, specified by
// itemId
//
// The variable `accessCount` in this function refers to the number of friend
// groups that the user is in that has access to the content item
func UserHasAccessToItem(username string, itemId int) bool {
	// This conditional must be at the top of the function so that users who are
	// not logged-in can still access the content item page of a public item
	if itemIsPublic(itemId) || userIsAuthor(username, itemId) {
		return true
	}
	row := db.QueryRow(`
	-- Get all the friend groups that the content item is shared in
	SELECT COUNT(fg_name) FROM Share
	WHERE item_id = ? AND fg_name IN (
    -- Get all the friend groups that the user belongs to
    SELECT fg_name FROM Belong
    WHERE member_email = ?
	)
	`, itemId, username)

	var accessCount int
	err := row.Scan(&accessCount)
	if err != nil {
		log.Println("content_item: UserHasAccessToItem() scan error: ", err)
	}

	return accessCount > 0
}

func itemIsPublic(itemId int) bool {
	row := db.QueryRow(`
	SELECT is_pub FROM Content_Item WHERE item_id = ?
	`, &itemId)

	var isPub bool
	err := row.Scan(&isPub)
	if err != nil {
		log.Println("content_item: itemIsPublic() scan error: ", err)
	}

	return isPub
}

func userIsAuthor(username string, itemId int) bool {
	row := db.QueryRow(`
	SELECT poster_email FROM Content_Item WHERE item_id = ?
	`, &itemId)

	var author string
	err := row.Scan(&author)
	if err != nil {
		log.Println("content_item: userIsAuthor() scan error: ", err)
	}

	return author == username
}
