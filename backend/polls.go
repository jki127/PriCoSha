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

// checkVote returns whether vote has already been cast
func checkVote(voterEmail string, itemID int) bool {
	row := db.QueryRow(`SELECT item_id
		FROM Vote
		WHERE voter_email=?
		AND item_id=?`,
		voterEmail, itemID)

	var temp int
	err := row.Scan(&temp)

	if err == nil {
		return true
	}
	return false
}

/*
AddVote takes a Vote primary key and add a row to the Vote table
*/
func AddVote(voterEmail string, itemID int, choice string) {
	// If user has voted in poll already, update choice
	if checkVote(voterEmail, itemID) {
		statement, err := db.Prepare(`UPDATE Vote
			SET choice=?
			WHERE voter_email=?
			AND item_id=?`)
		if err != nil {
			log.Println(`polls: AddVote(): Could not prepare update`)
		}

		_, err = statement.Exec(choice, voterEmail, itemID)
		if err != nil {
			log.Println(`polls: AddVote(): Could not execute update`)
		}
	} else {
		// Otherwise, add vote to poll
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
}
