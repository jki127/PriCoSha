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
}

// ContentItem holds info of Content_Item entities in the database
type ContentItem struct {
	ItemID   int
	Email    string
	FilePath string
	FileName string
	PostTime string // should use go date format later
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
		log.Println("Could not query public content from DB.")
	}
	defer rows.Close()

	// Declare variables for processing data
	var (
		itemID      int
		email       string
		filePath    string
		fileName    string
		postTime    string
		isPub       int
		data        []*ContentItem
		CurrentItem *ContentItem
	)

	for rows.Next() {
		err = rows.Scan(&itemID, &email, &filePath, &fileName, &postTime, &isPub)
		if err != nil {
			log.Println("Could not scan row data from public content query.")
		}
		CurrentItem = &ContentItem{
			ItemID:   itemID,
			Email:    email,
			FilePath: filePath,
			FileName: fileName,
			PostTime: postTime,
		}
		data = append(data, CurrentItem)
	}

	return data
}

func init() {
	var configData Conf

	configFile, err := os.Open("../backend/config.json")
	if err != nil {
		log.Println("Could not open config file.")
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println("Could not decode config file.")
	}

	dSN := configData.User + ":" + configData.Pass + "@/" + configData.DBName
	log.Println(dSN)

	db, err = sql.Open("mysql", dSN)
	if err != nil {
		log.Fatal("Could not connect to database")
	}
}
