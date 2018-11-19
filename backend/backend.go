// This is a placeholder backend .go file
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Used to interact with mySQL DB
	_ "github.com/go-sql-driver/mysql"
)

// Holds configuration for user data
type UserData struct {
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
	PostTime string 
}

var db *sql.DB

// TestDB tries to ping the database and returns the resulting error
func TestDB() error {
	return db.Ping()
}


func GetPubContent() []*ContentItem {
	
	rows, err := db.Query(`SELECT * FROM Content_Item 
		WHERE is_pub = true 
		AND post_time >= NOW() - INTERVAL 1 DAY`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Declare variables 
	var isPub bool
	var data[]*ContentItem
	var CurrentItem *ContentItem
	//iterate rows to add to array
	for rows.Next() {
		err = rows.Scan(&CurrentItem.ItemID, &CurrentItem.Email, 
			&CurrentItem.FilePath, &CurrentItem.FileName, &CurrentItem.PostTime, &isPub)
		if err != nil {
			panic(err)
		}
		data = append(data, CurrentItem)
	}
	return data
}

func ValidateInfo(username string, pass string) bool {
	
	rows, err := db.Query(`SELECT email FROM Person
		WHERE email=?
		AND password=SHA2(?,256)`,
		username, pass)
	if err != nil {
		log.Println("Failed Validation")
		reutrn false
	} else{
		reuturn true
	}

	

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

	dSN := configData.User + ":" + configData.Pass + "@/" + configData.DBName
	log.Println(dSN)

	db, err = sql.Open("mysql", dSN)
	if err != nil {
		log.Fatal("Could not connect to database")
	}
}
