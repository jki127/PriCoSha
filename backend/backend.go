// This is a placeholder backend .go file
package main

import (
	"log"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println(`I'm intended to interact with an sql server while handling
		requests from the frontend.`)
}
