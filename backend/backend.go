package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	// Will be used to interact with the mysql DB
	_ "github.com/go-sql-driver/mysql"
)

// Conf holds configuration data for DSN
type Conf struct {
	User   string
	Pass   string
	DBName string
}

var db *sql.DB

// TestDB tries to ping the database and returns the resulting error
func TestDB() error {
	return db.Ping()
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
