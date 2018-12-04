package backend

import (
	"log"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

func GetProfileData(username string) (fname string, lname string) {
	rows, err := db.Query(`SELECT f_name, l_name 
	FROM Person 
	WHERE email=?`,
		username)
	if err != nil {
		log.Println(`backend: getProfileData(): Could not
		query Profile Data from DB.`, username)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(fname, lname)
		if err != nil {
			log.Println(`backend: getProfileData(): Could not scan row data
			from Person content query.`)
		}
	}
	return fname, lname
}
