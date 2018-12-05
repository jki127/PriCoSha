package backend

import (
	"database/sql"
	"log"
)

/*
ValidateBelongFriendGroup takes an user's email and a FriendGroup primary key and
returns whether or not that email is found in the set of members of the FriendGroup
*/
func ValidateBelongFriendGroup(memberEmail string, fgName string,
	ownerEmail string) bool {
	row := db.QueryRow(`SELECT member_email FROM Belong
		WHERE member_email=?
		AND fg_name=? AND owner_email=?`,
		memberEmail, fgName, ownerEmail)

	var email string
	err := row.Scan(&email)

	switch {
	case err == sql.ErrNoRows:
		log.Println("backend: ValidateInfo(): no valid user found")
		return true
	case err != nil:
		log.Println("backend: ValidateInfo(): non nil Scan() error")
		return false
	default:
		return false
	}
}

/*
GetFriendGroup takes an user's email and returns a list of the FriendGroups that
the user owns
*/
func GetFriendGroup(userEmail string) []*FriendGroup {
	// Query DB for data
	rows, err := db.Query(`SELECT fg_name, owner_email, description 
		FROM Friend_Group 
		WHERE owner_email =?`, userEmail)
	if err != nil {
		log.Println(`backend: GetFriendGroup(): Could not
		query friend groups from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var FGData []*FriendGroup

	for rows.Next() {
		var CurrItem FriendGroup
		err = rows.Scan(&CurrItem.FGName, &CurrItem.OwnerEmail,
			&CurrItem.Description)
		if err != nil {
			log.Println(`backend: GetFriendGroup(): Could not scan row data
			from friend group query.`)
		}
		FGData = append(FGData, &CurrItem)
	}
	return FGData
}

/*
GetBelongFriendGroup takess the user's email and returns a list of Friend Groups
that the user belongs to (including own)
*/
func GetBelongFriendGroup(userEmail string) []*FriendGroup {
	// Query DB for data
	rows, err := db.Query(`SELECT fg_name, owner_email, description 
		FROM Friend_Group NATURAL JOIN Belong
		WHERE member_email =? AND owner_email !=?`, userEmail, userEmail)
	if err != nil {
		log.Println(`backend: GetBelongFriendGroup(): Could not
		query friend groups from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var BFGData []*FriendGroup

	for rows.Next() {
		var CurrItem FriendGroup
		err = rows.Scan(&CurrItem.FGName, &CurrItem.OwnerEmail,
			&CurrItem.Description)
		if err != nil {
			log.Println(`backend: GetBelongFriendGroup(): Could not scan row data
			from friend group query.`)
		}
		BFGData = append(BFGData, &CurrItem)
	}
	return BFGData
}
