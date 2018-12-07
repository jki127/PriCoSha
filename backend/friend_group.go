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
		log.Println(`friend_group: ValidateBelongFriendGroup(): User does not already 
			exist in Friend Group.`)
		return true
	case err != nil:
		log.Println("friend_group: ValidateBelongFriendGroup(): non nil Scan() error")
		return false
	default:
		log.Println(`friend_group: ValidateBelongFriendGroup(): User exists in 
			friend group.`)
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
BFGDataElement is designed to return the users role within the friend
group and the friendgroup info the role corresponds to
*/
type BFGDataElement struct {
	Role int
	FG   FriendGroup
}

/*
GetBelongFriendGroup takess the user's email and returns a list of Friend Groups
that the user belongs to and does not own
*/
func GetBelongFriendGroup(userEmail string) []*BFGDataElement {
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
	var BFGData []*BFGDataElement

	for rows.Next() {
		var CurrFG FriendGroup
		var CurrRole int
		err = rows.Scan(&CurrFG.FGName, &CurrFG.OwnerEmail,
			&CurrFG.Description)
		if err != nil {
			log.Println(`backend: GetBelongFriendGroup(): Could not scan row data
			from friend group query.`)
		}

		CurrRole = GetRole(CurrFG.FGName, CurrFG.OwnerEmail, userEmail)
		BFGData = append(BFGData, &BFGDataElement{
			Role: CurrRole,
			FG:   CurrFG,
		})
	}

	return BFGData
}
