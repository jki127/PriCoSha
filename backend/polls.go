package backend

import (
	"log"
)

/*
CheckPoll takes an itemID and returns true if the ContentItem is a poll
*/
func CheckPoll(itemID int) bool {
	row := db.QueryRow(`SELECT item_id
		FROM Content_Item
		WHERE format=1
		AND item_id=?`,
		itemID)

	var temp int
	err := row.Scan(&temp)

	switch {
	case err == nil:
		return true
	default:
		// if no results or error
		return false
	}
}

// Vote holds vote_count data from DB
type Vote struct {
	Choice string
	Count  int
}

/*
GetVotes takes an itemID and returns an array of vote structs
*/
func GetVotes(itemID int) []*Vote {
	rows, err := db.Query(`SELECT choice, COUNT(*) as vote_count
		FROM Vote
		WHERE item_id=?
		GROUP BY choice
		ORDER BY vote_count DESC`,
		itemID)
	if err != nil {
		log.Println(`polls: GetVotes(): Could not query DB`)
	}
	defer rows.Close()

	var data []*Vote

	for rows.Next() {
		var Current Vote

		err = rows.Scan(&Current.Choice, &Current.Count)
		if err != nil {
			log.Println(`polls: GetVotes(): Non-nil scan error`)
		}

		data = append(data, &Current)
	}

	return data
}

/*
AddVote takes a Vote primary key and add a row to the Vote table
*/
func AddVote(voterEmail string, itemID int, choice string) {
	statement, err := db.Prepare(`INSERT INTO Vote
		(voter_email, item_id, choice)
		VALUES
		(?, ?, ?)`)
	if err != nil {
		log.Println(`polls: AddVote(): Could not prepare insert`)
	}

	_, err = statement.Exec(voterEmail, itemID, choice)
	if err != nil {
		log.Println(`polls: AddVote(): Could not execute insert`)
	}
}
