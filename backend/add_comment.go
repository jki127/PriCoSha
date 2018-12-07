package backend

import (
	"log"
)

/*
ExecInsertComment prepares and executes statement to insert new comment into Comment
*/
func ExecInsertComment(comment Comment) {
	log.Println("add_comment: id:", comment.ItemID, "by", comment.Email)
	statement, err := db.Prepare(`INSERT INTO Comment VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Println(`add_comment: ExecInsertComment(): Could not prepare comment insertion`)
	}
	defer statement.Close()

	_, err = statement.Exec(comment.Email, comment.ItemID, comment.CommentTime, comment.Body)
	if err != nil {
		log.Println(`add_comment: ExecInsertComment(): Could not execute comment insertion`)
		log.Println(err)
	}
}
