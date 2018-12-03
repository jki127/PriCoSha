package main
import (
	"log"
	"html/template"
	"net/http"
	b "pricosha/backend"
)

func tagManagerHandler(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("username")
	var logged bool
	var username string
	if err != nil {
		logged = false
		username = ""
		log.Println("User was not logged in and cannot manage tags.")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		logged = true
		username = cookie.Value
	}

	CurrentTMD := TMD{
		Logged:	logged,
		Username: username,
		PendingTagData: b.GetPendingTags(username),
		AcceptedTagData: b.GetAcceptedTags(username),
	}
	t:=template.Must(template.ParseFiles("../web/template/tagmanager.html"))
	t.Execute(w, CurrentTMD)
}