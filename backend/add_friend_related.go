package backend

import "log"

/*
GetEmail takes a first name and a last name and returns the emails of users with
that name
*/
func GetEmail(fname string, lname string) []*string {

	//Checks for emails that have the same input name
	var emailList []*string
	rows, err := db.Query(`SELECT email FROM Person
				WHERE f_name=? 
				AND l_name=?`, fname, lname)
	if err != nil {
		log.Println(`add_friend_related: GetEmail(): Could not
		query Person's email from DB.`)
	}
	defer rows.Close()

	for rows.Next() {
		var PEmail string
		err = rows.Scan(&PEmail)
		if err != nil {
			log.Println(`add_friend_related: GetEmail(): Could not scan row data
			from emails query.`)
		}
		emailList = append(emailList, &PEmail)
	}

	return emailList
}

/*
AddFriend takes info of a friend entity and inserts the data into the Belong table
*/
func AddFriend(memberEmail string, fgname string, ownerEmail string) {
	statement, err := db.Prepare(`INSERT INTO Belong 
		(member_email, fg_name, owner_email)
		VALUES (?,?,?)`)
	if err != nil {
		log.Println(`add_friend_related: AddFriend(): Could not prepare insertion`)
	}
	defer statement.Close()
	_, err = statement.Exec(memberEmail, fgname, ownerEmail)

	if err != nil {
		log.Println(`add_friend_related: AddFriend(): Could not execute insertion`)
	}
}

func DeleteFriend(memberEmail string, fgname string, ownerEmail string){
	statement, err := db.Prepare(`DELETE FROM Belong WHERE member_email =? 
			AND fg_name =? 
			AND owner_email =?`)
	if err != nil {
		log.Println(`add_friend_related: DeleteFriend(): Could not prepare deletion`)
	}
	defer statement.Close()
	_, err = statement.Exec(memberEmail, fgname, ownerEmail)

	if err != nil {
		log.Println(`add_friend_related: DeleteFriend(): Could not execute deletion`)
	}
	log.Println ("Delete friend successfully!")
}