package backend

import "log"

/* This is a query meant to be used in the below functions instead of
retyping the same query over and over or creating a more complicated
SQL function. It uses a users email as TWO parameters so any
query that appends this string must pass two additional parameters
of the same variable in its execute phase. BE CAREFUL USING THIS. */
// Forgive me
var validItemSubQuery = `
SELECT item_id
FROM Content_Item
WHERE item_id IN (
	SELECT item_id 
	FROM Share
	WHERE (fg_name, owner_email) IN (
		SELECT fg_name, owner_email 
		FROM Belong
		WHERE member_email=?
	)
)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
OR poster_email=?`

/*
cleanTags checks for invalid tags and deletes them. Please see comment at
top of file before editing.
*/
func cleanTags(email string) {
	stmtStr := `
	DELETE FROM Tag
	WHERE item_id NOT IN (
		` + validItemSubQuery + `
	) 
	AND 
	(tagger_email=?
	OR
	tagged_email=?)`
	statement, err := db.Prepare(stmtStr)
	if err != nil {
		log.Println(`remove_hanging: cleanTags(): Could not prepare statement`)
		log.Println(err)
	}

	_, err = statement.Exec(email, email, email, email)
	if err != nil {
		log.Println(`remove_hanging: cleanTags(): Could not execute statement`)
	}
}

/*
cleanRates checks for invalid rates and deletes them. Please see comment at
top of file before editing.
*/
func cleanRates(email string) {
	stmtStr := `
	DELETE FROM Rate
	WHERE item_id NOT IN (
		` + validItemSubQuery + `
	) 
	AND email=?`
	statement, err := db.Prepare(stmtStr)
	if err != nil {
		log.Println(`remove_hanging: cleanRates(): Could not prepare statement`)
	}

	_, err = statement.Exec(email, email, email)
	if err != nil {
		log.Println(`remove_hanging: cleanRates(): Could not execute statement`)
	}
}

/*
cleanVotes checks for invalid votes and deletes them. Please see comment at
top of file before editing.
*/
func cleanVotes(email string) {
	stmtStr := `
	DELETE FROM Vote
	WHERE item_id NOT IN (
		` + validItemSubQuery + `
	) 
	AND voter_email=?`
	statement, err := db.Prepare(stmtStr)
	if err != nil {
		log.Println(`remove_hanging: cleanVotes(): Could not prepare statement`)
	}

	_, err = statement.Exec(email, email, email)
	if err != nil {
		log.Println(`remove_hanging: cleanVotes(): Could not execute statement`)
	}
}

/*
cleanComments checks for invalid comments and deletes them. Please see comment at
top of file before editing.
*/
func cleanComments(email string) {
	stmtStr := `
	DELETE FROM Comment
	WHERE item_id NOT IN (
		` + validItemSubQuery + `
	) 
	AND email=?`
	statement, err := db.Prepare(stmtStr)
	if err != nil {
		log.Println(`remove_hanging: cleanComments(): Could not prepare statement`)
	}

	_, err = statement.Exec(email, email, email)
	if err != nil {
		log.Println(`remove_hanging: cleanComments(): Could not execute statement`)
	}
}

func cleanInclude(email string) {
	stmtStr := `
	DELETE FROM Include
	WHERE item_id NOT IN (
		` + validItemSubQuery + `
	)
	AND email=?`
	statement, err := db.Prepare(stmtStr)
	if err != nil {
		log.Println(`remove_hanging: cleanInclude(): Could not prepare statement`)
	}
	_, err = statement.Exec(email, email, email)
	if err != nil {
		log.Println(`remove_hanging: cleanInclude(): Could not execute statement`)
	}
}

/*
CleanUp checks for invalid entities and deletes them.
*/
func CleanUp(email string) {
	log.Println("Cleaning...")
	// log.Println("1. Making view...")
	// makeView(email)
	log.Println("1. Cleaning tags...")
	cleanTags(email)
	log.Println("2. Cleaning rates...")
	cleanRates(email)
	log.Println("3. Cleaning votes...")
	cleanVotes(email)
	log.Println("4. Cleaning comments...")
	cleanComments(email)
	log.Println("5. Cleaning includes...")
	cleanInclude(email)
	log.Println("Cleaned up!")
}
