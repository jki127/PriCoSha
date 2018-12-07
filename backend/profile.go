package backend

import (
	"log"
	"math/rand"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

//GetProfileData queries DB for data of user
func GetProfileData(username string) (fname string, lname string, bio string, bioBool bool) {
	rows, err := db.Query(`SELECT f_name, l_name, bio
	FROM Person
	WHERE email=?`,
		username)
	if err != nil {
		log.Println(`backend: getProfileData(): Could not
		query Profile Data from DB.`, username)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&fname, &lname, &bio)
		if err != nil {
			log.Println(`backend: getProfileData(): Could not scan row data
			from Person content query.`, username)
		}
	}

	bioBool = false
	if bio != "" {
		bioBool = true
	}

	return fname, lname, bio, bioBool
}

//GetFriendsList queries the data base and returns an array of friends of the given user
func GetFriendsList(username string) []*FriendStruct {
	rows, err := db.Query(`SELECT DISTINCT member_email, f_name, l_name 
	FROM Belong JOIN Person ON Belong.member_email=Person.email
	WHERE Belong.owner_email=? AND Belong.member_email!=?`,
		username, username)

	// Declare variables for processing data
	var (
		data []*FriendStruct
	)

	defer rows.Close()

	for rows.Next() {
		var CurrentFriend FriendStruct
		err = rows.Scan(&CurrentFriend.FriendUsername, &CurrentFriend.FriendFirstName, &CurrentFriend.FriendLastName)
		if err != nil {
			log.Println(`backend: GetFriendsList(): Could not scan row data
				from friends in user's friendgroups query.`)
		}
		CurrentFriend.FaceID = rand.Intn(25)
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
			CurrentFriend.FaceID = rand.Intn(25)
			data = append(data, &CurrentFriend)
		}
	}

	return data
}

//AddBioToDB executes a query to the DB to add a bio to the user's entry
func AddBioToDB(username string, bio string) {
	_, err := db.Exec(`UPDATE Person SET bio=? WHERE email=?`, bio, username)
	if err != nil {
		log.Println(`backend: addBioToDB(): Could not
		Update bio in DB.`)
		return
	}
	log.Println("Bio added successfully")
}
