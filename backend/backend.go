package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

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

// FriendGroup holds info of Friend_Group entities in the database
type FriendGroup struct {
	MemberEmail string
	FGName      string
	OwnerEmail  string
	Description string
}

// Tag holds info of Tag entities in the database data
type Tag struct {
	TaggerEmail string
	TaggedEmail string
	ItemID      int
	TagTime     time.Time
	Status      bool // false = private; true = public
	FileName    string
	FilePath    string
}

//FriendStruct holds info on a Friend
type FriendStruct struct {
	FriendFirstName string
	FriendLastName  string
	FriendUsername  string
	FaceID          int
}
// ContentItem holds info related to ContentItem management
type ContentItem struct {
	ItemID   int
	Email    string
	FilePath string
	FileName string
	PostTime time.Time
	Fname    string 
	Lname    string
	RandImg  int
	Comments []*Comment
	Ratings  []*Rating
	IsPoll   bool    // 0 = not poll, 1 = poll
	Votes    []*Vote // holds related votes if
}

// Rating holds info related to Rating management
type Rating struct {
	Email     string
	Rate_time time.Time // WHO DID THIS // NO UNDERSCORES
	Emoji     string
}

// Comment holds info related to Comment management
type Comment struct {
	ItemID      int
	Email       string
	Body        string
	CommentTime time.Time
	Fname       string
	Lname       string
}

var db *sql.DB

// TestDB tries to ping the database and returns the resulting error
func TestDB() error {
	return db.Ping()
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

func init() {
	var configData Conf

	configFile, err := os.Open("../backend/config.json")
	if err != nil {
		log.Println("backend: init(): Could not open config file.")
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Println("backend: init(): Could not decode config file.")
	}

	if configData.Port != "" {
		configData.Port = ":" + configData.Port
	}

	dSN := configData.User + ":" + configData.Pass + "@(localhost" +
		configData.Port + ")/" + configData.DBName + "?parseTime=true" +
		"&charset=utf8mb4&collation=utf8mb4_unicode_ci"

	db, err = sql.Open("mysql", dSN)
	if err != nil {
		log.Fatal("backend: init(): Could not connect to database")
	}
}
