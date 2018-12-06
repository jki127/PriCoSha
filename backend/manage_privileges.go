package backend

import (
	"database/sql"
	"log"
)

/*
GetAtRole takes a FriendGroup primary key and a role int and returns the
users emails in that group with that role status
*/
func GetAtRole(fgName string, ownerEmail string, role int) []*string {
	rows, err := db.Query(`SELECT member_email
		FROM Belong
		WHERE fg_name=?
		AND owner_email=?
		AND role=?`,
		fgName, ownerEmail, role)
	if err != nil {
		log.Println(`manage_privileges: GetRoles(): Could not query DB.`)
	}
	defer rows.Close()

	var data []*string

	for rows.Next() {
		var current string
		err = rows.Scan(&current)
		if err != nil {
			log.Println(`manage_privileges: GetRoles(): Could now scan row
				data`)
		}
		data = append(data, &current)
	}
	return data
}

/*
GetRole takes a FriendGroup primary key and a member_email string and returns
the role status of the member
*/
func GetRole(fgName string, ownerEmail string, memberEmail string) int {
	log.Println("GetRoleData:", fgName, ownerEmail, memberEmail)
	row := db.QueryRow(`SELECT role
		FROM Belong
		WHERE fg_name=?
		AND owner_email=?
		AND member_email=?`,
		fgName, ownerEmail, memberEmail)

	var role int
	err := row.Scan(&role)

	switch {
	case err == sql.ErrNoRows:
		log.Println("manage_privileges: GetRole(): no valid user found")
		return -1
	case err != nil:
		log.Println("manage_privileges: GetRole(): non nil Scan() error")
		return -2
	default:
		return role
	}
}

/*
ChangePrivilege takes a Belong primary key and demotes or promotes the
associated member depending on the type of action (0 = demote, 1 = promote)
*/
func ChangePrivilege(fgName string, ownerEmail string, memberEmail string,
	action int) {
	statement, err := db.Prepare(`UPDATE Belong
		SET role=?
		WHERE member_email=?
		AND fg_name=?
		AND owner_email=?`)
	if err != nil {
		log.Println(`manage_privileges: ChangePrivilege(): Could not prepare
			update`)
	}
	defer statement.Close()

	var execErr error
	switch action {
	case 0:
		_, execErr = statement.Exec(2, memberEmail, fgName, ownerEmail)
	case 1:
		_, execErr = statement.Exec(1, memberEmail, fgName, ownerEmail)
	}

	if execErr != nil {
		log.Println(`manage_privileges: ChangePrivilege(): Could not execute
			update`)
		log.Println(err)
	}
}

/*
UserHasRemoveRights takes a FriendGroup primary key and a Share primary key
*/
func UserHasRemoveRights(fgName string, ownerEmail string, memberEmail string,
	itemID int) bool {
	log.Println("UserHasRemoveRights", fgName, ownerEmail)
	removes := UserCanRemoveFrom(memberEmail, itemID)
	for i := range removes {
		log.Println(removes[i])
		if removes[i].FGName == fgName && removes[i].OwnerEmail == ownerEmail {
			return true
		}
	}
	return false
}

/*
UserCanRemoveFrom takes a user's email and an item_id and returns a list of the
Friend_Groups that the user can unshare that item from
*/
func UserCanRemoveFrom(memberEmail string, itemID int) []*FriendGroup {
	/* Returns all groups the user has privileges for unsharing
	the Content_Item */
	rows, err := db.Query(`
	SELECT fg_name, owner_email
	FROM Share NATURAL JOIN Friend_Group NATURAL JOIN Belong
	-- Check if the user is the original poster of the Content_Item
	WHERE (member_email IN (
		SELECT poster_email
		FROM Content_Item
		WHERE item_id=?
	)
	-- Or if the user has mod privileges over the Shared Content_Item
	OR role < 2)
	AND member_email =?
	AND item_id = ?`,
		itemID, memberEmail, itemID)
	if err != nil {
		log.Println(`manage_privileges: UserCanRemoveFrom(): Could not query 
				DB.`)
	}
	defer rows.Close()

	var data []*FriendGroup

	for rows.Next() {
		var Current FriendGroup
		err = rows.Scan(&Current.FGName, &Current.OwnerEmail)
		if err != nil {
			log.Println(`manage_privileges: UserCanRemoveForm(): Could now scan row
				data`)
		}
		data = append(data, &Current)
	}

	return data
}

/*
UnshareItem takes a Share primary key and deletes the entry
*/
func UnshareItem(fgName string, ownerEmail string, itemID int) {
	statement, err := db.Prepare(`DELETE FROM Share
		WHERE fg_name=?
		AND owner_email=?
		AND item_id=?`)
	if err != nil {
		log.Println(`manage_privileges: UnshareItem(): Could not prepare
			delete`)
	}
	defer statement.Close()

	_, err = statement.Exec(fgName, ownerEmail, itemID)
	if err != nil {
		log.Println(`manage_privileges: UnshareItem(): Could not execute
			delete`)
		log.Println(err)
	}
}

/*
RenameFG takes a FriendGroup primary key and a string and renames it to that
string's value
*/
func RenameFG(fgName string, ownerEmail string, newName string) {
	// Get the description of the Friend_Group
	row := db.QueryRow(`SELECT description
		FROM Friend_Group
		WHERE fg_name=?
		AND owner_email=?`,
		fgName, ownerEmail)

	var description string
	err := row.Scan(&description)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not select group
			description`)
	}

	// Insert new Friend_Group with updated name
	insStatement, err := db.Prepare(`INSERT INTO Friend_Group
		(fg_name, owner_email, description)
		VALUES
		(?, ?, ?)`)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not prepare insert`)
	}
	defer insStatement.Close()

	// Update all rows in Belong with new name
	belongStatement, err := db.Prepare(`UPDATE Belong
		SET fg_name=?
		WHERE fg_name=?
		AND owner_email=?`)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not prepare belong 
			update`)
	}
	defer belongStatement.Close()

	// Update all rows in Share with new name
	shareStatement, err := db.Prepare(`UPDATE Share
		SET fg_name=?
		WHERE fg_name=?
		AND owner_email=?`)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not prepare share 
			update`)
	}
	defer shareStatement.Close()

	// Delete old Friend_Group with outdated name
	delStatement, err := db.Prepare(`DELETE FROM Friend_Group
		WHERE fg_name=?
		AND owner_email=?`)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not prepare delete`)
	}
	defer delStatement.Close()

	_, err = insStatement.Exec(newName, ownerEmail, description)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not execute insert`)
	}

	_, err = belongStatement.Exec(newName, fgName, ownerEmail)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not execute belong
			update`)
	}

	_, err = shareStatement.Exec(newName, fgName, ownerEmail)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not execute share
			update`)
	}

	_, err = delStatement.Exec(fgName, ownerEmail)
	if err != nil {
		log.Println(`manage_privileges: RenameFG(): Could not execute delete`)
	}
}
