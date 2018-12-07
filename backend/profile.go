package backend

import (
	"log"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

//GetProfileData queries DB for data of user
func GetProfileData(username string) (fname string, lname string) {
	// func GetProfileData(username string) (fname string, lname string, bio string) {
	rows, err := db.Query(`SELECT f_name, l_name 
	FROM Person 
	WHERE email=?`,
		username)
	// rows, err := db.Query(`SELECT f_name, l_name, bio
	// FROM Person
	// WHERE email=?`,
	// 	username)
	if err != nil {
		log.Println(`backend: getProfileData(): Could not
		query Profile Data from DB.`, username)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&fname, &lname)
		// err := rows.Scan(&fname, &lname, &bio)
		if err != nil {
			log.Println(`backend: getProfileData(): Could not scan row data
			from Person content query.`, username)
		}
	}
	return fname, lname
	// return fname, lname, bio
}

type FriendStruct struct {
	FriendFirstName string
	FriendLastName  string
	FriendUsername  string
}

func GetFriendsList(username string) []*FriendStruct {
	rows, err := db.Query(`SELECT DISTINCT member_email, f_name, l_name 
	FROM Belong JOIN Person ON Belong.member_email=Person.email
	WHERE Belong.owner_email=? AND Belong.member_email!=?`,
		username, username)

	// Declare variables for processing data
	var (
		data []*FriendStruct
	)
	for rows.Next() {
		var CurrentFriend FriendStruct
		err = rows.Scan(&CurrentFriend.FriendUsername, &CurrentFriend.FriendFirstName, &CurrentFriend.FriendLastName)
		if err != nil {
			log.Println(`backend: GetFriendsList(): Could not scan row data
				from friends in user's friendgroups query.`)
		}
		data = append(data, &CurrentFriend)
	}

	rows, err = db.Query(`SELECT DISTINCT owner_email, f_name, l_name 
	FROM Belong JOIN Person ON Belong.owner_email=Person.email
	WHERE member_email=? AND owner_email!=?`,
		username, username)

	for rows.Next() {
		var CurrentFriend FriendStruct
		err = rows.Scan(&CurrentFriend.FriendUsername, &CurrentFriend.FriendFirstName, &CurrentFriend.FriendLastName)
		if err != nil {
			log.Println(`backend: GetFriendsList(): Could not scan row data
					from friends in user's friendgroups query.`)
		}

		isPresent := false
		for _, aFriend := range data {
			if aFriend.FriendUsername == CurrentFriend.FriendUsername {
				isPresent = true
			}
		}
		if !isPresent {
			data = append(data, &CurrentFriend)
		}
	}

	return data
}
