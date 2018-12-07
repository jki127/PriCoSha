package main

import (
	"html/template"
	"net/http"
	b "pricosha/backend"
)

func contentLocationHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	urlParams := r.URL.Query()
	location := urlParams["loc"][0]

	pageData := struct {
		Logged       bool
		Username     string
		Location     string
		ContentItems []*b.ContentItem
	}{
		logged,
		username,
		location,
		b.GetUserContentByLocation(username, location),
	}

	t := template.Must(template.New("").ParseFiles("../web/template/content_location.html", "../web/template/base.html"))
	t.ExecuteTemplate(w, "base", pageData)
}
