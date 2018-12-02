package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

// Conf holds configuration data for DSN
type Conf struct {
	User   string
	Pass   string
	DBName string
	Port   string
}

// ContentItem holds info of Content_Item entities in the database
type ContentItem struct {
	ItemID   int
	Email    string
	FilePath string
	FileName string
	PostTime string // should use go date format later
}

// PendingTag holds info of Pending Tag items, taken from database data
type PendingTag struct {
	TaggerEmail string
	TaggedEmail string
	ItemID int
	TagTime string
	FileName string
	FilePath string
}

var db *sql.DB

// TestDB tries to ping the database and returns the resulting error
func TestDB() error {
	return db.Ping()
}

/*
GetPubContent queries DB for all Content_Item entities with a public
status and returns them as an array of ContentItem pointers.
*/
func GetPubContent() []*ContentItem {
	// Query DB for data
	rows, err := db.Query(`SELECT * FROM Content_Item
		WHERE is_pub = true
		AND post_time >= DATE_SUB(NOW(), INTERVAL 24 HOUR)`)
	if err != nil {
		log.Println(`backend: GetPubContent(): Could not
		query public content from DB.`)
	}
	defer rows.Close()

	// Declare variables for processing data
	var (
		isPub int
		data  []*ContentItem
	)

	for rows.Next() {
		var CurrentItem ContentItem
		err = rows.Scan(&CurrentItem.ItemID, &CurrentItem.Email,
			&CurrentItem.FilePath, &CurrentItem.FileName,
			&CurrentItem.PostTime, &isPub)
		if err != nil {
			log.Println(`backend: GetPubContent(): Could not scan row data
			from public content query.`)
		}
		data = append(data, &CurrentItem)
	}

	return data
}

func GetPendingTags(username string) []*PendingTag{
	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id, tag_time, file_name, file_path FROM Tag NATURAL JOIN Content_Item
		WHERE status = false
		AND tagged_email=?`,username)
	if err != nil {
		log.Println(`backend: GetPendingTags(): Could not
		query tags content from DB.`)
	}
	defer rows.Close()
	
	var data []*PendingTag

	for rows.Next() {
		var CurrentTag PendingTag
		err = rows.Scan(&CurrentTag.TaggerEmail, &CurrentTag.TaggedEmail,
			&CurrentTag.ItemID, &CurrentTag.TagTime,
			&CurrentTag.FileName, &CurrentTag.FilePath)
		if err != nil {
			log.Println(`backend: GetPendingTags(): Could not scan row data
			from tag content query.`)
		}
		data = append(data, &CurrentTag)
	}

	return data
}


/*
ValidateInfo receives user entered login info and queries the DB on whether
or not that info is valid
*/
func ValidateInfo(username string, password string) bool {
	// Query DB for data
	row := db.QueryRow(`SELECT email FROM Person
		WHERE email=?
		AND password=SHA2(?,256)`,
		username, password)

	var email string
	err := row.Scan(&email)

	switch {
	case err == sql.ErrNoRows:
		log.Println("backend: ValidateInfo(): no valid user found")
		return false
	case err != nil:
		log.Println("backend: ValidateInfo(): non nil Scan() error")
		return false
	default:
		return true
	}
}

// //Checks for Pending Tags on  a User
// func CheckPendingTags(taggedEmail string) ([]*string, []*int) {
// 	//Check for Tags with the User's email with status flagged as false
// 	var PendingTagsTagger []*string
// 	var PendingTagsItemID []*int
// 	rows, err := db.Query(`SELECT tagger_email, tagged_email, item_id FROM Tag WHERE tagged_email=?`,taggedEmail)
// 	if err!=nil{
// 		log.Println(`backend: CheckPendingTags(): Could not query Person's Pending Tags from DB.`)
// 	}
// 	defer rows.Close()

// 	for rows.Next(){
// 		var TaggerEmail string
// 		var ItemID int
// 		err = rows.Scan(&TaggerEmail, &ItemID)
// 		if err != nil {
// 			log.Println(`backend: CheckPendingTags(): Could not scan row data from Check Pending Tags query.`)
// 		}
// 		PendingTagsItemID=append(PendingTagsItemID,&ItemID)
// 		PendingTagsTagger=append(PendingTagsTagger,&TaggerEmail)

// 	}
// 	return PendingTagsTagger, PendingTagsItemID
// }




func init() {
	var configData Conf

	configFile, err := os.Open("../backend/config.json")
	if err != nil {
		log.Println("backend: init(): Could not open config file.")
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println("backend: init(): Could not decode config file.")
	}

	if configData.Port != "" {
		configData.Port = ":" + configData.Port
	}

	dSN := configData.User + ":" + configData.Pass + "@(localhost" +
		configData.Port + ")/" + configData.DBName

	db, err = sql.Open("mysql", dSN)
	if err != nil {
		log.Fatal("backend: init(): Could not connect to database")
	}
}
