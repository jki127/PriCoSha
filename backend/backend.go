// Package backend provides functions for connecting to the database as well as
// apply logic to collections derived from the database
package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

// UserData holds configuration for user data
type UserData struct {
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
	PostTime string
}

var db *sql.DB

// TestDB tries to ping the database and returns the resulting error
func TestDB() error {
	return db.Ping()
}

// GetPubContent retrieves all ContentItems from the database with is_pub set
// to true
func GetPubContent() []*ContentItem {
	rows, err := db.Query(`SELECT * FROM Content_Item
		WHERE is_pub = true
		AND post_time >= NOW() - INTERVAL 1 DAY`)
	if err != nil {
		log.Println("Database Error")
	}
	defer rows.Close()

	var data []*ContentItem

	//iterate rows to add to array
	for rows.Next() {
		var (
			isPub       bool
			currentItem ContentItem
		)

		err = rows.Scan(&currentItem.ItemID, &currentItem.Email,
			&currentItem.FilePath, &currentItem.FileName, &currentItem.PostTime, &isPub)
		if err != nil {
			log.Println("No Items Available")
		}

		data = append(data, &currentItem)
	}

	return data
}

// ValidateInfo validates the login credentials of a user
func ValidateInfo(username string, pass string) bool {
	//rows, err := db.Query(`SELECT email FROM Person
	_, err := db.Query(`SELECT email FROM Person
	WHERE email=?
	AND password=SHA2(?,256)`,
		username, pass)
	if err != nil {
		log.Println("Failed Validation")
		return false
	}
	return true
}

func init() {
	var configData UserData

	configFile, err := os.Open("../backend/config.json")
	if err != nil {
		log.Println("Could not open config file.")
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println("Could not decode config file.")
	}

	if configData.Port != "" {
		configData.Port = ":" + configData.Port
	}

	dSN := configData.User + ":" + configData.Pass +
		"@(localhost" + configData.Port + ")" + "/" + configData.DBName
	log.Println(dSN)

	db, err = sql.Open("mysql", dSN)
	if err != nil {
		log.Fatal("Could not connect to database")
	}
}
