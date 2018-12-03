package main
import (
	// "log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

func declineTagHandler(w http.ResponseWriter, r *http.Request){
	// cookie, err := r.Cookie("username")
	// // var logged bool
	// var username string
	// if err != nil {
	// 	// logged = false
	// 	username = ""
	// 	log.Println("User was not logged in and cannot decline tags.")
	// 	http.Redirect(w, r, "/login", http.StatusFound)
	// 	return
	// } else {
	// 	// logged = true
	// 	username = cookie.Value
	// }


	url := r.URL
	queryData :=url.Query()
	id, _:=strconv.Atoi(queryData["iid"][0])

	// if username==queryData["ted"][0] {
	// 	b.DeclineTag(queryData["ter"][0],queryData["ted"][0],id)
	// }
	b.DeclineTag(queryData["ter"][0],queryData["ted"][0],id)

	http.Redirect(w, r, "/login", http.StatusFound)
}