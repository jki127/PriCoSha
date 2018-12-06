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
